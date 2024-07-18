# VK_EXT_conservative_rasterization(3) Manual Page

## Name

VK_EXT_conservative_rasterization - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

102

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_fragment_fully_covered](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_fragment_fully_covered.html)

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_conservative_rasterization%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_conservative_rasterization%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-06-09

**Interactions and External Dependencies**  
- This extension requires
  [`SPV_EXT_fragment_fully_covered`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_fragment_fully_covered.html)
  if the
  `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`fullyCoveredFragmentShaderInputVariable`
  feature is used.

- This extension requires
  [`SPV_KHR_post_depth_coverage`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_post_depth_coverage.html)if
  the
  `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`conservativeRasterizationPostDepthCoverage`
  feature is used.

- This extension provides API support for
  [`GL_NV_conservative_raster_underestimation`](https://registry.khronos.org/OpenGL/extensions/NV/NV_conservative_raster_underestimation.txt)
  if the
  `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`fullyCoveredFragmentShaderInputVariable`
  feature is used.

**Contributors**  
- Daniel Koch, NVIDIA

- Daniel Rakos, AMD

- Jeff Bolz, NVIDIA

- Slawomir Grajewski, Intel

- Stu Smith, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

This extension adds a new rasterization mode called conservative
rasterization. There are two modes of conservative rasterization;
overestimation and underestimation.

When overestimation is enabled, if any part of the primitive, including
its edges, covers any part of the rectangular pixel area, including its
sides, then a fragment is generated with all coverage samples turned on.
This extension allows for some variation in implementations by
accounting for differences in overestimation, where the generating
primitive size is increased at each of its edges by some sub-pixel
amount to further increase conservative pixel coverage. Implementations
can allow the application to specify an extra overestimation beyond the
base overestimation the implementation already does. It also allows
implementations to either cull degenerate primitives or rasterize them.

When underestimation is enabled, fragments are only generated if the
rectangular pixel area is fully covered by the generating primitive. If
supported by the implementation, when a pixel rectangle is fully covered
the fragment shader input variable builtin called FullyCoveredEXT is set
to true. The shader variable works in either overestimation or
underestimation mode.

Implementations can process degenerate triangles and lines by either
discarding them or generating conservative fragments for them.
Degenerate triangles are those that end up with zero area after the
rasterizer quantizes them to the fixed-point pixel grid. Degenerate
lines are those with zero length after quantization.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceConservativeRasterizationPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceConservativeRasterizationPropertiesEXT.html)

- Extending
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html):

  - [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConservativeRasterizationModeEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkPipelineRasterizationConservativeStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_CONSERVATIVE_RASTERIZATION_EXTENSION_NAME`

- `VK_EXT_CONSERVATIVE_RASTERIZATION_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT`

## <a href="#_new_built_in_variables" class="anchor"></a>New Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-fullycoveredext"
  target="_blank" rel="noopener">FullyCoveredEXT</a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-FragmentFullyCoveredEXT"
  target="_blank" rel="noopener">FragmentFullyCoveredEXT</a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1.1, 2020-09-06 (Piers Daniell)

  - Add missing SPIR-V and GLSL dependencies.

- Revision 1, 2017-08-28 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_conservative_rasterization"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
