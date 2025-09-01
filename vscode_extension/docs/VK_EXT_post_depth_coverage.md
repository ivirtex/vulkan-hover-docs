# VK\_EXT\_post\_depth\_coverage(3) Manual Page

## Name

VK\_EXT\_post\_depth\_coverage - device extension



## [](#_registered_extension_number)Registered Extension Number

156

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_post\_depth\_coverage](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_post_depth_coverage.html)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_post_depth_coverage%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_post_depth_coverage%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-07-17

**Interactions and External Dependencies**

- This extension provides API support for [`GL_ARB_post_depth_coverage`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_post_depth_coverage.txt) and [`GL_EXT_post_depth_coverage`](https://registry.khronos.org/OpenGL/extensions/EXT/EXT_post_depth_coverage.txt)

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_KHR_post_depth_coverage`

which allows the fragment shader to control whether values in the `SampleMask` built-in input variable reflect the coverage after early [depth](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth) and [stencil](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-stencil) tests are applied.

This extension adds a new `PostDepthCoverage` execution mode under the `SampleMaskPostDepthCoverage` capability. When this mode is specified along with `EarlyFragmentTests`, the value of an input variable decorated with the [`SampleMask`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-samplemask) built-in reflects the coverage after the early fragment tests are applied. Otherwise, it reflects the coverage before the depth and stencil tests.

When using GLSL source-based shading languages, the `post_depth_coverage` layout qualifier from GL\_ARB\_post\_depth\_coverage or GL\_EXT\_post\_depth\_coverage maps to the `PostDepthCoverage` execution mode.

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_POST_DEPTH_COVERAGE_EXTENSION_NAME`
- `VK_EXT_POST_DEPTH_COVERAGE_SPEC_VERSION`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`SampleMaskPostDepthCoverage`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-SampleMaskPostDepthCoverage)

## [](#_version_history)Version History

- Revision 1, 2017-07-17 (Daniel Koch)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_post_depth_coverage).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0