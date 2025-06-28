# VkVideoPictureResourceInfoKHR(3) Manual Page

## Name

VkVideoPictureResourceInfoKHR - Structure specifying the parameters of a video picture resource



## [](#_c_specification)C Specification

The `VkVideoPictureResourceInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_video_queue
typedef struct VkVideoPictureResourceInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkOffset2D         codedOffset;
    VkExtent2D         codedExtent;
    uint32_t           baseArrayLayer;
    VkImageView        imageViewBinding;
} VkVideoPictureResourceInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `codedOffset` is the offset in texels of the image subregion to use.
- `codedExtent` is the size in pixels of the coded image data.
- `baseArrayLayer` is the array layer of the image view specified in `imageViewBinding` to use as the video picture resource.
- `imageViewBinding` is an image view representing the video picture resource.

## [](#_description)Description

The image subresource referred to by such a structure is defined as the image array layer index specified in `baseArrayLayer` relative to the image subresource range the image view specified in `imageViewBinding` was created with.

The meaning of the `codedOffset` and `codedExtent` depends on the command and context the video picture resource is used in, as well as on the used [video profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles) and corresponding codec-specific semantics, as described later.

A video picture resource is uniquely defined by the image subresource referred to by an instance of this structure, together with the `codedOffset` and `codedExtent` members that identify the image subregion within the image subresource referenced corresponding to the video picture resource according to the particular codec-specific semantics.

Accesses to image data within a video picture resource happen at the granularity indicated by [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`pictureAccessGranularity`, as returned by [vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html) for the used [video profile](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-profiles). As a result, given an effective image subregion corresponding to a video picture resource, the actual image subregion accessed **may** be larger than that as it **may** include additional padding texels due to the picture access granularity. Any writes performed by video coding operations to such padding texels will result in undefined texel values.

Two video picture resources match if they refer to the same image subresource and they specify identical `codedOffset` and `codedExtent` values.

Valid Usage

- [](#VUID-VkVideoPictureResourceInfoKHR-baseArrayLayer-07175)VUID-VkVideoPictureResourceInfoKHR-baseArrayLayer-07175  
  `baseArrayLayer` **must** be less than the [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`subresourceRange.layerCount` specified when the image view `imageViewBinding` was created

Valid Usage (Implicit)

- [](#VUID-VkVideoPictureResourceInfoKHR-sType-sType)VUID-VkVideoPictureResourceInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VIDEO_PICTURE_RESOURCE_INFO_KHR`
- [](#VUID-VkVideoPictureResourceInfoKHR-pNext-pNext)VUID-VkVideoPictureResourceInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkVideoPictureResourceInfoKHR-imageViewBinding-parameter)VUID-VkVideoPictureResourceInfoKHR-imageViewBinding-parameter  
  `imageViewBinding` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkOffset2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoDecodeInfoKHR.html), [VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeInfoKHR.html), [VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoReferenceSlotInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoPictureResourceInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0