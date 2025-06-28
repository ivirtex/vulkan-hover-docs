# VK\_NV\_sample\_mask\_override\_coverage(3) Manual Page

## Name

VK\_NV\_sample\_mask\_override\_coverage - device extension



## [](#_registered_extension_number)Registered Extension Number

95

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_sample\_mask\_override\_coverage](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_sample_mask_override_coverage.html)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_sample_mask_override_coverage%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_sample_mask_override_coverage%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-12-08

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_sample_mask_override_coverage`](https://registry.khronos.org/OpenGL/extensions/NV/NV_sample_mask_override_coverage.txt)

**Contributors**

- Daniel Koch, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_NV_sample_mask_override_coverage`

The extension provides access to the `OverrideCoverageNV` decoration under the `SampleMaskOverrideCoverageNV` capability. Adding this decoration to a variable with the `SampleMask` builtin decoration allows the shader to modify the coverage mask and affect which samples are used to process the fragment.

When using GLSL source-based shader languages, the `override_coverage` layout qualifier from `GL_NV_sample_mask_override_coverage` maps to the `OverrideCoverageNV` decoration. To use the `override_coverage` layout qualifier in GLSL the `GL_NV_sample_mask_override_coverage` extension must be enabled. Behavior is described in the `GL_NV_sample_mask_override_coverage` extension spec.

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_SAMPLE_MASK_OVERRIDE_COVERAGE_EXTENSION_NAME`
- `VK_NV_SAMPLE_MASK_OVERRIDE_COVERAGE_SPEC_VERSION`

## [](#_new_variable_decoration)New Variable Decoration

- [OverrideCoverageNV in SampleMask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-samplemask)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`SampleMaskOverrideCoverageNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-SampleMaskOverrideCoverageNV)

## [](#_version_history)Version History

- Revision 1, 2016-12-08 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_sample_mask_override_coverage)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0