# VK\_EXT\_depth\_clip\_enable(3) Manual Page

## Name

VK\_EXT\_depth\_clip\_enable - device extension



## [](#_registered_extension_number)Registered Extension Number

103

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_clip_enable%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_clip_enable%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-12-20

**Contributors**

- Daniel Rakos, AMD
- Henri Verbeet, CodeWeavers
- Jeff Bolz, NVIDIA
- Philip Rebohle, DXVK
- Tobias Hector, AMD

## [](#_description)Description

This extension allows the depth clipping operation, that is normally implicitly controlled by [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html)::`depthClampEnable`, to instead be controlled explicitly by [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html)::`depthClipEnable`.

This is useful for translating DX content which assumes depth clamping is always enabled, but depth clip can be controlled by the DepthClipEnable rasterization state (D3D12\_RASTERIZER\_DESC).

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDepthClipEnableFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthClipEnableFeaturesEXT.html)
- Extending [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html):
  
  - [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPipelineRasterizationDepthClipStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationDepthClipStateCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEPTH_CLIP_ENABLE_EXTENSION_NAME`
- `VK_EXT_DEPTH_CLIP_ENABLE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLIP_ENABLE_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_DEPTH_CLIP_STATE_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 1, 2018-12-20 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_depth_clip_enable)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0