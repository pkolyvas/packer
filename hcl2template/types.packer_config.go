package hcl2template

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/packer/packer"
)

// PackerConfig represents a loaded packer config
type PackerConfig struct {
	Sources map[SourceRef]*Source

	Variables PackerV1Variables

	Builds Builds
}

func (p *Parser) CoreBuildProvisioners(blocks []*ProvisionerBlock, generatedVars []string) ([]packer.CoreBuildProvisioner, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	res := []packer.CoreBuildProvisioner{}
	for _, pb := range blocks {
		provisioner, moreDiags := p.StartProvisioner(pb, generatedVars)
		diags = append(diags, moreDiags...)
		if moreDiags.HasErrors() {
			continue
		}
		res = append(res, packer.CoreBuildProvisioner{
			PType:       pb.PType,
			PName:       pb.PName,
			Provisioner: provisioner,
		})
	}
	return res, diags
}

func (p *Parser) CoreBuildPostProcessors(blocks []*PostProcessorBlock) ([]packer.CoreBuildPostProcessor, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	res := []packer.CoreBuildPostProcessor{}
	for _, ppb := range blocks {
		postProcessor, moreDiags := p.StartPostProcessor(ppb)
		diags = append(diags, moreDiags...)
		if moreDiags.HasErrors() {
			continue
		}
		res = append(res, packer.CoreBuildPostProcessor{
			PostProcessor: postProcessor,
			PName:         ppb.PName,
			PType:         ppb.PType,
		})
	}

	return res, diags
}

func (p *Parser) getBuilds(cfg *PackerConfig) ([]packer.Build, hcl.Diagnostics) {
	res := []packer.Build{}
	var diags hcl.Diagnostics

	for _, build := range cfg.Builds {
		for _, from := range build.Froms {
			src, found := cfg.Sources[from]
			if !found {
				diags = append(diags, &hcl.Diagnostic{
					Summary:  "Unknown " + sourceLabel + " " + from.String(),
					Subject:  build.HCL2Ref.DefRange.Ptr(),
					Severity: hcl.DiagError,
				})
				continue
			}
			builder, moreDiags, generatedVars := p.StartBuilder(src)
			diags = append(diags, moreDiags...)
			if moreDiags.HasErrors() {
				continue
			}
			provisioners, moreDiags := p.CoreBuildProvisioners(build.ProvisionerBlocks, generatedVars)
			diags = append(diags, moreDiags...)
			if moreDiags.HasErrors() {
				continue
			}
			postProcessors, moreDiags := p.CoreBuildPostProcessors(build.PostProcessors)
			pps := [][]packer.CoreBuildPostProcessor{}
			if len(postProcessors) > 0 {
				pps = [][]packer.CoreBuildPostProcessor{postProcessors}
			}
			diags = append(diags, moreDiags...)
			if moreDiags.HasErrors() {
				continue
			}

			pcb := &packer.CoreBuild{
				Type:           src.Type,
				Builder:        builder,
				Provisioners:   provisioners,
				PostProcessors: pps,
				Variables:      cfg.Variables,
			}
			res = append(res, pcb)
		}
	}
	return res, diags
}

func (p *Parser) Parse(path string) ([]packer.Build, hcl.Diagnostics) {
	cfg, diags := p.parse(path)
	if diags.HasErrors() {
		return nil, diags
	}

	builds, moreDiags := p.getBuilds(cfg)
	return builds, append(diags, moreDiags...)
}
