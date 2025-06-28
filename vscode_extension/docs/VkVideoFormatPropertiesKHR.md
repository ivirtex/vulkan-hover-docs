# VkVideoFormatPropertiesKHR(3) Manual Page

## Name

VkVideoFormatPropertiesKHR - Structure enumerating the video image formats



## [](#_c_specification)C Specification

The `VkVideoFormatPropertiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoFormatPropertiesKHR {
    VkStructureType       sType;
    void*                 pNext;
    VkFormat              format;
    VkComponentMapping    componentMapping;
    VkImageCreateFlags    imageCreateFlags;
    VkImageType           imageType;
    VkImageTiling         imageTiling;
    VkImageUsageFlags     imageUsageFlags;
} VkVideoFormatPropertiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) that specifies the format that **can** be used with the specified video profiles and image usages.
- `componentMapping` defines the color channel order used for the format. `format` along with `componentMapping` describe how the color channels are ordered when producing video decoder output or are expected to be ordered in video encoder input, when applicable. If the `format` reported does not require component swizzling then all members of `componentMapping` will be set to `VK_COMPONENT_SWIZZLE_IDENTITY`.
- `imageCreateFlags` is a bitmask of [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html) specifying the supported image creation flags for the format.
- `imageType` is a [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html) that specifies the image type the format **can** be used with.
- `imageTiling` is a [VkImageTiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageTiling.html) that specifies the image tiling the format **can** be used with.
- `imageUsageFlags` is a bitmask of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) specifying the supported image usage flags for the format.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkVideoFormatPropertiesKHR-sType-sType)VUID-VkVideoFormatPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_FORMAT_PROPERTIES_KHR`
- [](#VUID-VkVideoFormatPropertiesKHR-pNext-pNext)VUID-VkVideoFormatPropertiesKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkVideoFormatAV1QuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatAV1QuantizationMapPropertiesKHR.html), [VkVideoFormatH265QuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatH265QuantizationMapPropertiesKHR.html), or [VkVideoFormatQuantizationMapPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoFormatQuantizationMapPropertiesKHR.html)
- [](#VUID-VkVideoFormatPropertiesKHR-sType-unique)VUID-VkVideoFormatPropertiesKHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlags.html), [VkImageTiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoFormatPropertiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0