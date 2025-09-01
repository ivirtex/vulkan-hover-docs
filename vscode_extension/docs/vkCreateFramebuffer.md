# vkCreateFramebuffer(3) Manual Page

## Name

vkCreateFramebuffer - Create a new framebuffer object



## [](#_c_specification)C Specification

To create a framebuffer, call:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateFramebuffer(
    VkDevice                                    device,
    const VkFramebufferCreateInfo*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkFramebuffer*                              pFramebuffer);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the framebuffer.
- `pCreateInfo` is a pointer to a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html) structure describing additional information about framebuffer creation.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pFramebuffer` is a pointer to a [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html) handle in which the resulting framebuffer object is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateFramebuffer-device-10002)VUID-vkCreateFramebuffer-device-10002  
  `device` **must** support at least one queue family with the `VK_QUEUE_GRAPHICS_BIT` capability
- [](#VUID-vkCreateFramebuffer-pCreateInfo-02777)VUID-vkCreateFramebuffer-pCreateInfo-02777  
  If `pCreateInfo->flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and `attachmentCount` is not `0`, each element of `pCreateInfo->pAttachments` **must** have been created on `device`

Valid Usage (Implicit)

- [](#VUID-vkCreateFramebuffer-device-parameter)VUID-vkCreateFramebuffer-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateFramebuffer-pCreateInfo-parameter)VUID-vkCreateFramebuffer-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html) structure
- [](#VUID-vkCreateFramebuffer-pAllocator-parameter)VUID-vkCreateFramebuffer-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateFramebuffer-pFramebuffer-parameter)VUID-vkCreateFramebuffer-pFramebuffer-parameter  
  `pFramebuffer` **must** be a valid pointer to a [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html) handle
- [](#VUID-vkCreateFramebuffer-device-queuecount)VUID-vkCreateFramebuffer-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html), [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateFramebuffer).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0