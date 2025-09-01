# VkImageView(3) Manual Page

## Name

VkImageView - Opaque handle to an image view object



## [](#_c_specification)C Specification

Image objects are not directly accessed by pipeline shaders for reading or writing image data. Instead, *image views* representing contiguous ranges of the image subresources and containing additional metadata are used for that purpose. Views **must** be created on images of compatible types, and **must** represent a valid subset of image subresources.

Image views are represented by `VkImageView` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkImageView)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html), [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html), [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html), [VkImageViewCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCaptureDescriptorDataInfoEXT.html), [VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewHandleInfoNVX.html), [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html), [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentInfo.html), [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html), [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html), [VkVideoEncodeQuantizationMapInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeQuantizationMapInfoKHR.html), [VkVideoPictureResourceInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoPictureResourceInfoKHR.html), [vkBindOpticalFlowSessionImageNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindOpticalFlowSessionImageNV.html), [vkCmdBindInvocationMaskHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindInvocationMaskHUAWEI.html), [vkCmdBindShadingRateImageNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindShadingRateImageNV.html), [vkCreateImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImageView.html), [vkDestroyImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyImageView.html), [vkGetImageViewAddressNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageViewAddressNVX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageView).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0