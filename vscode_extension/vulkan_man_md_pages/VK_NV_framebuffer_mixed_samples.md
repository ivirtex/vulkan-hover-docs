# VK_NV_framebuffer_mixed_samples(3) Manual Page

## Name

VK_NV_framebuffer_mixed_samples - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

153

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_framebuffer_mixed_samples%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_framebuffer_mixed_samples%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-06-04

**Contributors**  
- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows multisample rendering with a raster and
depth/stencil sample count that is larger than the color sample count.
Rasterization and the results of the depth and stencil tests together
determine the portion of a pixel that is “covered”. It can be useful to
evaluate coverage at a higher frequency than color samples are stored.
This coverage is then “reduced” to a collection of covered color
samples, each having an opacity value corresponding to the fraction of
the color sample covered. The opacity can optionally be blended into
individual color samples.

Rendering with fewer color samples than depth/stencil samples greatly
reduces the amount of memory and bandwidth consumed by the color buffer.
However, converting the coverage values into opacity introduces
artifacts where triangles share edges and **may** not be suitable for
normal triangle mesh rendering.

One expected use case for this functionality is Stencil-then-Cover path
rendering (similar to the OpenGL GL_NV_path_rendering extension). The
stencil step determines the coverage (in the stencil buffer) for an
entire path at the higher sample frequency, and then the cover step
draws the path into the lower frequency color buffer using the coverage
information to antialias path edges. With this two-step process,
internal edges are fully covered when antialiasing is applied and there
is no corruption on these edges.

The key features of this extension are:

- It allows render pass and framebuffer objects to be created where the
  number of samples in the depth/stencil attachment in a subpass is a
  multiple of the number of samples in the color attachments in the
  subpass.

- A coverage reduction step is added to Fragment Operations which
  converts a set of covered raster/depth/stencil samples to a set of
  color samples that perform blending and color writes. The coverage
  reduction step also includes an optional coverage modulation step,
  multiplying color values by a fractional opacity corresponding to the
  number of associated raster/depth/stencil samples covered.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html):

  - [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageModulationModeNV.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkPipelineCoverageModulationStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateFlagsNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_FRAMEBUFFER_MIXED_SAMPLES_EXTENSION_NAME`

- `VK_NV_FRAMEBUFFER_MIXED_SAMPLES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-06-04 (Jeff Bolz)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_framebuffer_mixed_samples"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
