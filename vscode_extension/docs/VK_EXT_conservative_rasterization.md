# VK\_EXT\_conservative\_rasterization(3) Manual Page

## Name

VK\_EXT\_conservative\_rasterization - device extension



## [](#_registered_extension_number)Registered Extension Number

102

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_fragment\_fully\_covered](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_fragment_fully_covered.html)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_conservative_rasterization%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_conservative_rasterization%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-06-09

**Interactions and External Dependencies**

- This extension requires [`SPV_EXT_fragment_fully_covered`](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_fragment_fully_covered.html) if the `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`fullyCoveredFragmentShaderInputVariable` feature is used.
- This extension requires [`SPV_KHR_post_depth_coverage`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_post_depth_coverage.html)if the `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`conservativeRasterizationPostDepthCoverage` feature is used.
- This extension provides API support for [`GL_NV_conservative_raster_underestimation`](https://registry.khronos.org/OpenGL/extensions/NV/NV_conservative_raster_underestimation.txt) if the `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`fullyCoveredFragmentShaderInputVariable` feature is used.

**Contributors**

- Daniel Koch, NVIDIA
- Daniel Rakos, AMD
- Jeff Bolz, NVIDIA
- Slawomir Grajewski, Intel
- Stu Smith, Imagination Technologies

## [](#_description)Description

This extension adds a new rasterization mode called conservative rasterization. There are two modes of conservative rasterization; overestimation and underestimation.

When overestimation is enabled, if any part of the primitive, including its edges, covers any part of the rectangular pixel area, including its sides, then a fragment is generated with all coverage samples turned on. This extension allows for some variation in implementations by accounting for differences in overestimation, where the generating primitive size is increased at each of its edges by some sub-pixel amount to further increase conservative pixel coverage. Implementations can allow the application to specify an extra overestimation beyond the base overestimation the implementation already does. It also allows implementations to either cull degenerate primitives or rasterize them.

When underestimation is enabled, fragments are only generated if the rectangular pixel area is fully covered by the generating primitive. If supported by the implementation, when a pixel rectangle is fully covered the fragment shader input variable builtin called FullyCoveredEXT is set to true. The shader variable works in either overestimation or underestimation mode.

Implementations can process degenerate triangles and lines by either discarding them or generating conservative fragments for them. Degenerate triangles are those that end up with zero area after the rasterizer quantizes them to the fixed-point pixel grid. Degenerate lines are those with zero length after quantization.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceConservativeRasterizationPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceConservativeRasterizationPropertiesEXT.html)
- Extending [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html):
  
  - [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConservativeRasterizationModeEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPipelineRasterizationConservativeStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationConservativeStateCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_CONSERVATIVE_RASTERIZATION_EXTENSION_NAME`
- `VK_EXT_CONSERVATIVE_RASTERIZATION_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CONSERVATIVE_RASTERIZATION_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT`

## [](#_new_built_in_variables)New Built-In Variables

- [FullyCoveredEXT](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-fullycoveredext)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [FragmentFullyCoveredEXT](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-FragmentFullyCoveredEXT)

## [](#_version_history)Version History

- Revision 1.1, 2020-09-06 (Piers Daniell)
  
  - Add missing SPIR-V and GLSL dependencies.
- Revision 1, 2017-08-28 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_conservative_rasterization).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0