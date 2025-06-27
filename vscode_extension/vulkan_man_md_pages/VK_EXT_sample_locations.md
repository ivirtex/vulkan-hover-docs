# VK_EXT_sample_locations(3) Manual Page

## Name

VK_EXT_sample_locations - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

144

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_sample_locations%5D%20@drakos-amd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_sample_locations%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>drakos-amd</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to modify the locations of samples
within a pixel used in rasterization. Additionally, it allows
applications to specify different sample locations for each pixel in a
group of adjacent pixels, which **can** increase antialiasing quality
(particularly if a custom resolve shader is used that takes advantage of
these different locations).

It is common for implementations to optimize the storage of depth values
by storing values that **can** be used to reconstruct depth at each
sample location, rather than storing separate depth values for each
sample. For example, the depth values from a single triangle **may** be
represented using plane equations. When the depth value for a sample is
needed, it is automatically evaluated at the sample location. Modifying
the sample locations causes the reconstruction to no longer evaluate the
same depth values as when the samples were originally generated, thus
the depth aspect of a depth/stencil attachment **must** be cleared
before rendering to it using different sample locations.

Some implementations **may** need to evaluate depth image values while
performing image layout transitions. To accommodate this, instances of
the [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html) structure
**can** be specified for each situation where an explicit or automatic
layout transition has to take place.
[VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html) **can** be
chained from [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html)
structures to provide sample locations for layout transitions performed
by [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html) and
[vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html) calls, and
[VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html)
**can** be chained from
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html) to provide sample
locations for layout transitions performed implicitly by a render pass
instance.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEXT.html)

- [vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkAttachmentSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleLocationsEXT.html)

- [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html)

- [VkSampleLocationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationEXT.html)

- [VkSubpassSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassSampleLocationsEXT.html)

- Extending [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier.html),
  [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryBarrier2.html):

  - [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)

- Extending
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html):

  - [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)

- Extending [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html):

  - [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SAMPLE_LOCATIONS_EXTENSION_NAME`

- `VK_EXT_SAMPLE_LOCATIONS_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_MULTISAMPLE_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_SAMPLE_LOCATIONS_STATE_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_SAMPLE_LOCATIONS_BEGIN_INFO_EXT`

  - `VK_STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-08-02 (Daniel Rakos)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_sample_locations"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
