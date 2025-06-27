# VK_EXT_post_depth_coverage(3) Manual Page

## Name

VK_EXT_post_depth_coverage - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

156

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_post_depth_coverage](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_post_depth_coverage.html)

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_post_depth_coverage%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_post_depth_coverage%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-07-17

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_ARB_post_depth_coverage`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_post_depth_coverage.txt)
  and
  [`GL_EXT_post_depth_coverage`](https://registry.khronos.org/OpenGL/extensions/EXT/EXT_post_depth_coverage.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the following SPIR-V extension in
Vulkan:

- `SPV_KHR_post_depth_coverage`

which allows the fragment shader to control whether values in the
`SampleMask` built-in input variable reflect the coverage after early <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth"
target="_blank" rel="noopener">depth</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-stencil"
target="_blank" rel="noopener">stencil</a> tests are applied.

This extension adds a new `PostDepthCoverage` execution mode under the
`SampleMaskPostDepthCoverage` capability. When this mode is specified
along with `EarlyFragmentTests`, the value of an input variable
decorated with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-samplemask"
target="_blank" rel="noopener"><code>SampleMask</code></a> built-in
reflects the coverage after the early fragment tests are applied.
Otherwise, it reflects the coverage before the depth and stencil tests.

When using GLSL source-based shading languages, the
`post_depth_coverage` layout qualifier from GL_ARB_post_depth_coverage
or GL_EXT_post_depth_coverage maps to the `PostDepthCoverage` execution
mode.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_POST_DEPTH_COVERAGE_EXTENSION_NAME`

- `VK_EXT_POST_DEPTH_COVERAGE_SPEC_VERSION`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-SampleMaskPostDepthCoverage"
  target="_blank"
  rel="noopener"><code>SampleMaskPostDepthCoverage</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-07-17 (Daniel Koch)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_post_depth_coverage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
