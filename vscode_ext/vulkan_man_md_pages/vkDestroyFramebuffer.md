# vkDestroyFramebuffer(3) Manual Page

## Name

vkDestroyFramebuffer - Destroy a framebuffer object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a framebuffer, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyFramebuffer(
    VkDevice                                    device,
    VkFramebuffer                               framebuffer,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the framebuffer.

- `framebuffer` is the handle of the framebuffer to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyFramebuffer-framebuffer-00892"
  id="VUID-vkDestroyFramebuffer-framebuffer-00892"></a>
  VUID-vkDestroyFramebuffer-framebuffer-00892  
  All submitted commands that refer to `framebuffer` **must** have
  completed execution

- <a href="#VUID-vkDestroyFramebuffer-framebuffer-00893"
  id="VUID-vkDestroyFramebuffer-framebuffer-00893"></a>
  VUID-vkDestroyFramebuffer-framebuffer-00893  
  If `VkAllocationCallbacks` were provided when `framebuffer` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyFramebuffer-framebuffer-00894"
  id="VUID-vkDestroyFramebuffer-framebuffer-00894"></a>
  VUID-vkDestroyFramebuffer-framebuffer-00894  
  If no `VkAllocationCallbacks` were provided when `framebuffer` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyFramebuffer-device-parameter"
  id="VUID-vkDestroyFramebuffer-device-parameter"></a>
  VUID-vkDestroyFramebuffer-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyFramebuffer-framebuffer-parameter"
  id="VUID-vkDestroyFramebuffer-framebuffer-parameter"></a>
  VUID-vkDestroyFramebuffer-framebuffer-parameter  
  If `framebuffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `framebuffer` **must** be a valid [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html)
  handle

- <a href="#VUID-vkDestroyFramebuffer-pAllocator-parameter"
  id="VUID-vkDestroyFramebuffer-pAllocator-parameter"></a>
  VUID-vkDestroyFramebuffer-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyFramebuffer-framebuffer-parent"
  id="VUID-vkDestroyFramebuffer-framebuffer-parent"></a>
  VUID-vkDestroyFramebuffer-framebuffer-parent  
  If `framebuffer` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `framebuffer` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyFramebuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
