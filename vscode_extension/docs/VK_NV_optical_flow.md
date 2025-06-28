# VK\_NV\_optical\_flow(3) Manual Page

## Name

VK\_NV\_optical\_flow - device extension



## [](#_registered_extension_number)Registered Extension Number

465

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html)  
     and  
     [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_contact)Contact

- Carsten Rohde [\[GitHub\]crohde](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_optical_flow%5D%20%40crohde%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_optical_flow%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-09-26

**Contributors**

- Carsten Rohde, NVIDIA
- Vipul Parashar, NVIDIA
- Jeff Bolz, NVIDIA
- Eric Werness, NVIDIA

## [](#_description)Description

Optical flow are fundamental algorithms in computer vision (CV) area. This extension allows applications to estimate 2D displacement of pixels between two frames.

Note

This extension is designed to be used with upcoming NVIDIA Optical Flow SDK Version 5 which will be available on NVIDIA Developer webpage.

## [](#_new_object_types)New Object Types

- [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html)

## [](#_new_commands)New Commands

- [vkBindOpticalFlowSessionImageNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindOpticalFlowSessionImageNV.html)
- [vkCmdOpticalFlowExecuteNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdOpticalFlowExecuteNV.html)
- [vkCreateOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateOpticalFlowSessionNV.html)
- [vkDestroyOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyOpticalFlowSessionNV.html)
- [vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html)

## [](#_new_structures)New Structures

- [VkOpticalFlowExecuteInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteInfoNV.html)
- [VkOpticalFlowImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatPropertiesNV.html)
- [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html)
- Extending [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html):
  
  - [VkOpticalFlowSessionCreatePrivateDataInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreatePrivateDataInfoNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceOpticalFlowFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceOpticalFlowFeaturesNV.html)
- Extending [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html):
  
  - [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatInfoNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceOpticalFlowPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceOpticalFlowPropertiesNV.html)

## [](#_new_enums)New Enums

- [VkOpticalFlowExecuteFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteFlagBitsNV.html)
- [VkOpticalFlowGridSizeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagBitsNV.html)
- [VkOpticalFlowPerformanceLevelNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowPerformanceLevelNV.html)
- [VkOpticalFlowSessionBindingPointNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionBindingPointNV.html)
- [VkOpticalFlowSessionCreateFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateFlagBitsNV.html)
- [VkOpticalFlowUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowUsageFlagBitsNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkOpticalFlowExecuteFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowExecuteFlagsNV.html)
- [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowGridSizeFlagsNV.html)
- [VkOpticalFlowSessionCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateFlagsNV.html)
- [VkOpticalFlowUsageFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowUsageFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_OPTICAL_FLOW_EXTENSION_NAME`
- `VK_NV_OPTICAL_FLOW_SPEC_VERSION`
- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV`
  - `VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
  - `VK_FORMAT_R16G16_S10_5_NV`
  - `VK_FORMAT_R16G16_SFIXED5_NV`
- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_OPTICAL_FLOW_COST_BIT_NV`
  - `VK_FORMAT_FEATURE_2_OPTICAL_FLOW_IMAGE_BIT_NV`
  - `VK_FORMAT_FEATURE_2_OPTICAL_FLOW_VECTOR_BIT_NV`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_OPTICAL_FLOW_SESSION_NV`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV`
- Extending [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFlagBits.html):
  
  - `VK_QUEUE_OPTICAL_FLOW_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_OPTICAL_FLOW_EXECUTE_INFO_NV`
  - `VK_STRUCTURE_TYPE_OPTICAL_FLOW_IMAGE_FORMAT_INFO_NV`
  - `VK_STRUCTURE_TYPE_OPTICAL_FLOW_IMAGE_FORMAT_PROPERTIES_NV`
  - `VK_STRUCTURE_TYPE_OPTICAL_FLOW_SESSION_CREATE_INFO_NV`
  - `VK_STRUCTURE_TYPE_OPTICAL_FLOW_SESSION_CREATE_PRIVATE_DATA_INFO_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_OPTICAL_FLOW_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_OPTICAL_FLOW_PROPERTIES_NV`

## [](#_examples)Examples

```cpp
// Example querying available input formats
VkOpticalFlowImageFormatInfoNV ofFormatInfo = { VK_STRUCTURE_TYPE_OPTICAL_FLOW_IMAGE_FORMAT_INFO_NV };
ofFormatInfo.usage = VK_OPTICAL_FLOW_USAGE_INPUT_BIT_NV;

uint32_t count = 0;
vkGetPhysicalDeviceOpticalFlowImageFormatsNV(physicalDevice, &ofFormatInfo, &count, NULL);
VkOpticalFlowImageFormatPropertiesNV* fmt = new VkOpticalFlowImageFormatPropertiesNV[count];
memset(fmt, 0, count  * sizeof(VkOpticalFlowImageFormatPropertiesNV));
for (uint32_t i = 0; i < count; i++) {
    fmt[i].sType = VK_STRUCTURE_TYPE_OPTICAL_FLOW_IMAGE_FORMAT_PROPERTIES_NV;
}
vkGetPhysicalDeviceOpticalFlowImageFormatsNV(physicalDevice, &ofFormatInfo, &count, fmt);

// Pick one of the available formats
VkFormat inputFormat = fmt[0].format;

// Check feature support for optimal tiling
VkFormatProperties3 formatProperties3 = { VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_3 };
VkFormatProperties2 formatProperties2 = { VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2, &formatProperties3 };
vkGetPhysicalDeviceFormatProperties2(physicalDevice, inputFormat, &formatProperties2);
if (!(formatProperties3.optimalTilingFeatures & VK_FORMAT_FEATURE_2_OPTICAL_FLOW_IMAGE_BIT_NV)) {
    return false;
}

// Check support for image creation parameters
VkPhysicalDeviceImageFormatInfo2 imageFormatInfo2 = { VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2, &ofFormatInfo };
imageFormatInfo2.format = inputFormat;
imageFormatInfo2.type = VK_IMAGE_TYPE_2D;
imageFormatInfo2.tiling = VK_IMAGE_TILING_OPTIMAL;
imageFormatInfo2.usage = VK_IMAGE_USAGE_SAMPLED_BIT | VK_IMAGE_USAGE_TRANSFER_DST_BIT;

VkImageFormatProperties2 imageFormatProperties2 = { VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2 };
if (vkGetPhysicalDeviceImageFormatProperties2(physicalDevice, &imageFormatInfo2, &imageFormatProperties2) != VK_SUCCESS) {
    return false;
}

VkImageCreateInfo imageCreateInfo = { VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO, &ofFormatInfo };
imageCreateInfo.imageType = VK_IMAGE_TYPE_2D;
imageCreateInfo.format = inputFormat;
imageCreateInfo.extent = { width, height, (uint32_t)1};
imageCreateInfo.mipLevels = 1;
imageCreateInfo.arrayLayers = 1;
imageCreateInfo.samples = VK_SAMPLE_COUNT_1_BIT;
imageCreateInfo.usage = VK_IMAGE_USAGE_SAMPLED_BIT | VK_IMAGE_USAGE_TRANSFER_DST_BIT;
imageCreateInfo.tiling = VK_IMAGE_TILING_OPTIMAL;

vkCreateImage(device, &imageCreateInfo, NULL, &input);
"allocate memory, bind image, create view"

"do the same for reference and output"

// Create optical flow session
VkOpticalFlowSessionCreateInfoNV sessionCreateInfo = { VK_STRUCTURE_TYPE_OPTICAL_FLOW_SESSION_CREATE_INFO_NV };
sessionCreateInfo.width = width;
sessionCreateInfo.height = height;
sessionCreateInfo.imageFormat = inputFormat;
sessionCreateInfo.flowVectorFormat = outputFormat;
sessionCreateInfo.outputGridSize = VK_OPTICAL_FLOW_GRID_SIZE_4X4_BIT_NV;
sessionCreateInfo.performanceLevel = VK_OPTICAL_FLOW_PERFORMANCE_LEVEL_SLOW_NV;
VkOpticalFlowSessionNV session;
vkCreateOpticalFlowSessionNV(device, &sessionCreateInfo, NULL, &session);

"allocate command buffer"

"transfer images to VK_PIPELINE_STAGE_2_OPTICAL_FLOW_BIT_NV"
"transfer input images to VK_ACCESS_2_OPTICAL_FLOW_READ_BIT_NV and output image to VK_ACCESS_2_OPTICAL_FLOW_WRITE_BIT_NV"

vkBindOpticalFlowSessionImageNV(device, session, VK_OPTICAL_FLOW_SESSION_BINDING_POINT_INPUT_NV, inputView, VK_IMAGE_LAYOUT_GENERAL);
vkBindOpticalFlowSessionImageNV(device, session, VK_OPTICAL_FLOW_SESSION_BINDING_POINT_REFERENCE_NV, refView, VK_IMAGE_LAYOUT_GENERAL);
vkBindOpticalFlowSessionImageNV(device, session, VK_OPTICAL_FLOW_SESSION_BINDING_POINT_FLOW_VECTOR_NV, outputView, VK_IMAGE_LAYOUT_GENERAL);

VkOpticalFlowExecuteInfoNV opticalFlowExecuteInfo = { VK_STRUCTURE_TYPE_OPTICAL_FLOW_EXECUTE_INFO_NV };
vkCmdOpticalFlowExecuteNV(cmd, session, &opticalFlowExecuteInfo);

"submit command buffer"
```

## [](#_version_history)Version History

- Revision 1, 2022-09-26 (Carsten Rohde)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_optical_flow)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0