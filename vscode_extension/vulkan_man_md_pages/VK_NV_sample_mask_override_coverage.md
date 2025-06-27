# VK_NV_sample_mask_override_coverage(3) Manual Page

## Name

VK_NV_sample_mask_override_coverage - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

95

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_sample_mask_override_coverage](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_sample_mask_override_coverage.html)

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_sample_mask_override_coverage%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_sample_mask_override_coverage%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-12-08

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_sample_mask_override_coverage`](https://registry.khronos.org/OpenGL/extensions/NV/NV_sample_mask_override_coverage.txt)

**Contributors**  
- Daniel Koch, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_NV_sample_mask_override_coverage`

The extension provides access to the `OverrideCoverageNV` decoration
under the `SampleMaskOverrideCoverageNV` capability. Adding this
decoration to a variable with the `SampleMask` builtin decoration allows
the shader to modify the coverage mask and affect which samples are used
to process the fragment.

When using GLSL source-based shader languages, the `override_coverage`
layout qualifier from `GL_NV_sample_mask_override_coverage` maps to the
`OverrideCoverageNV` decoration. To use the `override_coverage` layout
qualifier in GLSL the `GL_NV_sample_mask_override_coverage` extension
must be enabled. Behavior is described in the
`GL_NV_sample_mask_override_coverage` extension spec.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_SAMPLE_MASK_OVERRIDE_COVERAGE_EXTENSION_NAME`

- `VK_NV_SAMPLE_MASK_OVERRIDE_COVERAGE_SPEC_VERSION`

## <a href="#_new_variable_decoration" class="anchor"></a>New Variable Decoration

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-samplemask"
  target="_blank" rel="noopener">OverrideCoverageNV in SampleMask</a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-SampleMaskOverrideCoverageNV"
  target="_blank"
  rel="noopener"><code>SampleMaskOverrideCoverageNV</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-12-08 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_sample_mask_override_coverage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
