# VkImage(3) Manual Page

## Name

VkImage - Opaque handle to an image object



## [](#_c_specification)C Specification

Images represent multidimensional - up to 3 - arrays of data which **can** be used for various purposes (e.g. attachments, textures), by binding them to a graphics or compute pipeline via descriptor sets, or by directly specifying them as parameters to certain commands.

Images are represented by `VkImage` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkImage)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html), [VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageInfo2.html), [VkCopyBufferToImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferToImageInfo2.html), [VkCopyImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageInfo2.html), [VkCopyImageToBufferInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToBufferInfo2.html), [VkCopyImageToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfo.html), [VkCopyImageToMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToMemoryInfo.html), [VkCopyMemoryToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToImageInfo.html), [VkDedicatedAllocationMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDedicatedAllocationMemoryAllocateInfoNV.html), [VkExportMetalIOSurfaceInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalIOSurfaceInfoEXT.html), [VkExportMetalTextureInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalTextureInfoEXT.html), [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFrameBoundaryEXT.html), [VkHostImageLayoutTransitionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageLayoutTransitionInfo.html), [VkImageCaptureDescriptorDataInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCaptureDescriptorDataInfoEXT.html), [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html), [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2.html), [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2.html), [VkImageSparseMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSparseMemoryRequirementsInfo2.html), [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html), [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfo.html), [VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveImageInfo2.html), [VkSparseImageMemoryBindInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryBindInfo.html), [VkSparseImageOpaqueMemoryBindInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageOpaqueMemoryBindInfo.html), [vkBindImageMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindImageMemory.html), [vkCmdBlitImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage.html), [vkCmdClearColorImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdClearColorImage.html), [vkCmdClearDepthStencilImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdClearDepthStencilImage.html), [vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage.html), [vkCmdCopyImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage.html), [vkCmdCopyImageToBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer.html), [vkCmdCopyMemoryToImageIndirectNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToImageIndirectNV.html), [vkCmdResolveImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResolveImage.html), [vkCreateImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImage.html), [vkDestroyImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyImage.html), [vkGetImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageDrmFormatModifierPropertiesEXT.html), [vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements.html), [vkGetImageSparseMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSparseMemoryRequirements.html), [vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout.html), [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html), [vkGetImageSubresourceLayout2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2EXT.html), [vkGetImageSubresourceLayout2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2KHR.html), [vkGetSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSwapchainImagesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImage)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0