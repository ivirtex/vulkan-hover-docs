# VK\_QCOM\_fragment\_density\_map\_offset(3) Manual Page

## Name

VK\_QCOM\_fragment\_density\_map\_offset - device extension



## [](#_registered_extension_number)Registered Extension Number

426

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_EXT\_fragment\_density\_map\_offset](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map_offset.html) extension

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_fragment_density_map_offset%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_fragment_density_map_offset%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-06-17

**Contributors**

- Matthew Netsch, Qualcomm Technologies, Inc.
- Jonathan Wicks, Qualcomm Technologies, Inc.
- Jonathan Tinkham, Qualcomm Technologies, Inc.
- Jeff Leger, Qualcomm Technologies, Inc.
- Manan Katwala, Qualcomm Technologies, Inc.
- Connor Abbott, Valve Corporation

## [](#_description)Description

This extension allows an application to specify offsets to a fragment density map attachment, changing the location where the fragment density map is applied to the framebuffer. This is helpful for eye-tracking use cases where the fovea needs to be moved around the framebuffer to match the eye pose.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM.html)
- Extending [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html), [VkRenderingEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingEndInfoEXT.html):
  
  - [VkSubpassFragmentDensityMapOffsetEndInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassFragmentDensityMapOffsetEndInfoQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_FRAGMENT_DENSITY_MAP_OFFSET_EXTENSION_NAME`
- `VK_QCOM_FRAGMENT_DENSITY_MAP_OFFSET_SPEC_VERSION`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_OFFSET_FEATURES_QCOM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_OFFSET_PROPERTIES_QCOM`
  - `VK_STRUCTURE_TYPE_SUBPASS_FRAGMENT_DENSITY_MAP_OFFSET_END_INFO_QCOM`

## [](#_examples)Examples

### [](#_fragment_density_map)Fragment Density Map

If the fragment density map size is larger than framebuffer size / `minFragmentDensityTexelSize`, then offsets may be applied with [VkSubpassFragmentDensityMapOffsetEndInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassFragmentDensityMapOffsetEndInfoQCOM.html) to shift the fragment density map.

```c++
// Create fragment density map
VkImageCreateInfo imageCreateInfo =
{
    .sType = VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO,
    .pNext = nullptr,
    .flags = VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM,
    .imageType = VK_IMAGE_TYPE_2D,    // Must be 2D
    .format = VK_FORMAT_R8G8_UNORM,   // Must have VK_FORMAT_FEATURE_FRAGMENT_DENSITY_MAP_BIT_EXT
    .extent = {64, 64, 1},
    .mipLevels = 1,
    .arrayLayers = 2,                 // 1 for each multiview view
    .samples = VK_SAMPLE_COUNT_1_BIT, // Must be 1x MSAA
    .tiling = tiling,
    .usage = VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT,
    // ...
};

vkCreateImage(device, &imageCreateInfo, nullptr, &fdmImage);

// Start render pass with fbo that has fdmImage as fragmentDensityMapAttachment
// All other attachments must have been created with
// VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM
// FBO has dimensions 1024x1024 in this example
vkCmdBeginRenderPass2(commandBuffer, &renderPassBeginInfo, pSubpassBeginInfo);

// ...

// Shift fragment density map image by offset to put fovea in center of vision
/*
   If minFragmentDensityTexelSize is 32x32, then offsets can be at most
   1024x1024 for this example before density value access is clamped to edge of map

   This is because the region of each fdmImage texel is

     texelSize = max(framebuffer / fdmImage, minFragmentDensityTexelSize) = 32x32

   and the fdmImage covers a total framebuffer region of

     fdmImageRegion = fdmImage * texelSize = 2048x2048

   This means that the fragment density map can be shifted up to 1024x1024
   before the density values are clamped to edge of map
*/
VkSubpassFragmentDensityMapOffsetEndInfoQCOM offsetInfo =
{
    .sType = VK_STRUCTURE_TYPE_SUBPASS_FRAGMENT_DENSITY_MAP_OFFSET_END_INFO_QCOM,
    .fragmentDensityOffsetCount = 2,    // 1 for each layer/multiview view
    .pFragmentDensityOffsets = offsets, // aligned to fragmentDensityOffsetGranularity
};

subpassEndInfo.pNext = &offsetInfo;

// Only offsets given to the last subpass are used for the whole render pass
// Offsets given in other subpasses are ignored
vkCmdEndRenderPass2(VkCommandBuffer commandBuffer, &subpassEndInfo);
```

## [](#_version_history)Version History

- Revision 3, 2025-03-20 (Connor Abbott/Matthew Netsch)
  
  - Fix clamp equation and clarify the offsets are FDM relative (editorial only)
- Revision 2, 2024-06-17 (Matthew Netsch)
  
  - Fix typo in spec regarding fragmentDensityMapOffset feature
- Revision 1, 2021-09-03 (Matthew Netsch)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_fragment_density_map_offset)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0