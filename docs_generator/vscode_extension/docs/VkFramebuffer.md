# VkFramebuffer(3) Manual Page

## Name

VkFramebuffer - Opaque handle to a framebuffer object



## [](#_c_specification)C Specification

Render passes operate in conjunction with *framebuffers*. Framebuffers represent a collection of specific memory attachments that a render pass instance uses.

Framebuffers are represented by `VkFramebuffer` handles:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkFramebuffer)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html), [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html), [vkCreateFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateFramebuffer.html), [vkDestroyFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyFramebuffer.html), [vkGetFramebufferTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFramebufferTilePropertiesQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFramebuffer)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0