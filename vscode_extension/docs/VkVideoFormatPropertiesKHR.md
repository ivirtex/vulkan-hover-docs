# VkVideoFormatPropertiesKHR(3) Manual Page

## Name

VkVideoFormatPropertiesKHR - Structure enumerating the video image
formats



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoFormatPropertiesKHR` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) that specifies the format that
  **can** be used with the specified video profiles and image usages.

- `componentMapping` defines the color channel order used for the
  format. `format` along with `componentMapping` describe how the color
  channels are ordered when producing video decoder output or are
  expected to be ordered in video encoder input, when applicable. If the
  `format` reported does not require component swizzling then all
  members of `componentMapping` will be set to
  `VK_COMPONENT_SWIZZLE_IDENTITY`.

- `imageCreateFlags` is a bitmask of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html) specifying the
  supported image creation flags for the format.

- `imageType` is a [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) that specifies the
  image type the format **can** be used with.

- `imageTiling` is a [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) that specifies
  the image tiling the format **can** be used with.

- `imageUsageFlags` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) specifying the
  supported image usage flags for the format.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkVideoFormatPropertiesKHR-sType-sType"
  id="VUID-VkVideoFormatPropertiesKHR-sType-sType"></a>
  VUID-VkVideoFormatPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_FORMAT_PROPERTIES_KHR`

- <a href="#VUID-VkVideoFormatPropertiesKHR-pNext-pNext"
  id="VUID-VkVideoFormatPropertiesKHR-pNext-pNext"></a>
  VUID-VkVideoFormatPropertiesKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlags.html),
[VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceVideoFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoFormatPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoFormatPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
