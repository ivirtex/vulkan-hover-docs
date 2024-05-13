# vkCreateFramebuffer(3) Manual Page

## Name

vkCreateFramebuffer - Create a new framebuffer object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a framebuffer, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateFramebuffer(
    VkDevice                                    device,
    const VkFramebufferCreateInfo*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkFramebuffer*                              pFramebuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the framebuffer.

- `pCreateInfo` is a pointer to a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html) structure
  describing additional information about framebuffer creation.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pFramebuffer` is a pointer to a [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html)
  handle in which the resulting framebuffer object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreateFramebuffer-pCreateInfo-02777"
  id="VUID-vkCreateFramebuffer-pCreateInfo-02777"></a>
  VUID-vkCreateFramebuffer-pCreateInfo-02777  
  If `pCreateInfo->flags` does not include
  `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and `attachmentCount` is not
  `0`, each element of `pCreateInfo->pAttachments` **must** have been
  created on `device`

Valid Usage (Implicit)

- <a href="#VUID-vkCreateFramebuffer-device-parameter"
  id="VUID-vkCreateFramebuffer-device-parameter"></a>
  VUID-vkCreateFramebuffer-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateFramebuffer-pCreateInfo-parameter"
  id="VUID-vkCreateFramebuffer-pCreateInfo-parameter"></a>
  VUID-vkCreateFramebuffer-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html) structure

- <a href="#VUID-vkCreateFramebuffer-pAllocator-parameter"
  id="VUID-vkCreateFramebuffer-pAllocator-parameter"></a>
  VUID-vkCreateFramebuffer-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateFramebuffer-pFramebuffer-parameter"
  id="VUID-vkCreateFramebuffer-pFramebuffer-parameter"></a>
  VUID-vkCreateFramebuffer-pFramebuffer-parameter  
  `pFramebuffer` **must** be a valid pointer to a
  [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html),
[VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateFramebuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
