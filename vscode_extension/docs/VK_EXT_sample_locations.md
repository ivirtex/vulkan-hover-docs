# VK\_EXT\_sample\_locations(3) Manual Page

## Name

VK\_EXT\_sample\_locations - device extension



## [](#_registered_extension_number)Registered Extension Number

144

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]drakos-amd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_sample_locations%5D%20%40drakos-amd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_sample_locations%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-08-02

**Contributors**

- Mais Alnasser, AMD
- Matthaeus G. Chajdas, AMD
- Maciej Jesionowski, AMD
- Daniel Rakos, AMD
- Slawomir Grajewski, Intel
- Jeff Bolz, NVIDIA
- Bill Licea-Kane, Qualcomm

## [](#_description)Description

This extension allows an application to modify the locations of samples within a pixel used in rasterization. Additionally, it allows applications to specify different sample locations for each pixel in a group of adjacent pixels, which **can** increase antialiasing quality (particularly if a custom resolve shader is used that takes advantage of these different locations).

It is common for implementations to optimize the storage of depth values by storing values that **can** be used to reconstruct depth at each sample location, rather than storing separate depth values for each sample. For example, the depth values from a single triangle **may** be represented using plane equations. When the depth value for a sample is needed, it is automatically evaluated at the sample location. Modifying the sample locations causes the reconstruction to no longer evaluate the same depth values as when the samples were originally generated, thus the depth aspect of a depth/stencil attachment **must** be cleared before rendering to it using different sample locations.

Some implementations **may** need to evaluate depth image values while performing image layout transitions. To accommodate this, instances of the [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html) structure **can** be specified for each situation where an explicit or automatic layout transition has to take place. [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html) **can** be chained from [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html) structures to provide sample locations for layout transitions performed by [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html) and [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html) calls, and [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html) **can** be chained from [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html) to provide sample locations for layout transitions performed implicitly by a render pass instance.

## [](#_new_commands)New Commands

- [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetSampleLocationsEXT.html)
- [vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html)

## [](#_new_structures)New Structures

- [VkAttachmentSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentSampleLocationsEXT.html)
- [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisamplePropertiesEXT.html)
- [VkSampleLocationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationEXT.html)
- [VkSubpassSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassSampleLocationsEXT.html)
- Extending [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html), [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2.html):
  
  - [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)
- Extending [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html):
  
  - [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)
- Extending [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html):
  
  - [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SAMPLE_LOCATIONS_EXTENSION_NAME`
- `VK_EXT_SAMPLE_LOCATIONS_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT`
  - `VK_STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT`

## [](#_issues)Issues

1\) When using with Dynamic Rendering, is there a VkRenderPassSampleLocationsBeginInfoEXT equivalent struct

**RESOLVED**: No, there are no subpasses that need to have a sample location set.

## [](#_version_history)Version History

- Revision 1, 2017-08-02 (Daniel Rakos)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_sample_locations)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0