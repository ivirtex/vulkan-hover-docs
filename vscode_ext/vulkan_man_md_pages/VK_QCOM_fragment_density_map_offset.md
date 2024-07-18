# VK_QCOM_fragment_density_map_offset(3) Manual Page

## Name

VK_QCOM_fragment_density_map_offset - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

426

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  
and  
[VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_fragment_density_map_offset%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_fragment_density_map_offset%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-06-17

**Contributors**  
- Matthew Netsch, Qualcomm Technologies, Inc.

- Jonathan Wicks, Qualcomm Technologies, Inc.

- Jonathan Tinkham, Qualcomm Technologies, Inc.

- Jeff Leger, Qualcomm Technologies, Inc.

- Manan Katwala, Qualcomm Technologies, Inc.

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to specify offsets to a fragment
density map attachment, changing the framebuffer location where density
values are applied to without having to regenerate the fragment density
map.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM.html)

- Extending [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassEndInfo.html):

  - [VkSubpassFragmentDensityMapOffsetEndInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassFragmentDensityMapOffsetEndInfoQCOM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_FRAGMENT_DENSITY_MAP_OFFSET_EXTENSION_NAME`

- `VK_QCOM_FRAGMENT_DENSITY_MAP_OFFSET_SPEC_VERSION`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_OFFSET_FEATURES_QCOM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_OFFSET_PROPERTIES_QCOM`

  - `VK_STRUCTURE_TYPE_SUBPASS_FRAGMENT_DENSITY_MAP_OFFSET_END_INFO_QCOM`

## <a href="#_examples" class="anchor"></a>Examples

### <a href="#_fragment_density_map" class="anchor"></a>Fragment Density Map

If the fragment density map size is larger than framebuffer size /
`minFragmentDensityTexelSize`, then offsets may be applied with
[VkSubpassFragmentDensityMapOffsetEndInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassFragmentDensityMapOffsetEndInfoQCOM.html)
to shift the framebuffer inside of the density map image. This is
helpful if apps want to reuse the same density map image to track
content without editing the image.

``` c
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

   This means that the framebuffer (1024x1024) can be shifted up to 1024x1024
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2024-06-17 (Matthew Netsch)

  - Fix typo in spec regarding fragmentDensityMapOffset feature

- Revision 1, 2021-09-03 (Matthew Netsch)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_fragment_density_map_offset"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
