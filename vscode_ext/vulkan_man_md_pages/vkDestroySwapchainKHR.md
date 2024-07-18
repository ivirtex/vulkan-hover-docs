# vkDestroySwapchainKHR(3) Manual Page

## Name

vkDestroySwapchainKHR - Destroy a swapchain object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a swapchain object call:

``` c
// Provided by VK_KHR_swapchain
void vkDestroySwapchainKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) associated with `swapchain`.

- `swapchain` is the swapchain to destroy.

- `pAllocator` is the allocator used for host memory allocated for the
  swapchain object when there is no more specific allocator available
  (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a>).

## <a href="#_description" class="anchor"></a>Description

The application **must** not destroy a swapchain until after completion
of all outstanding operations on images that were acquired from the
swapchain. `swapchain` and all associated `VkImage` handles are
destroyed, and **must** not be acquired or used any more by the
application. The memory of each `VkImage` will only be freed after that
image is no longer used by the presentation engine. For example, if one
image of the swapchain is being displayed in a window, the memory for
that image **may** not be freed until the window is destroyed, or
another swapchain is created for the window. Destroying the swapchain
does not invalidate the parent `VkSurfaceKHR`, and a new swapchain
**can** be created with it.

When a swapchain associated with a display surface is destroyed, if the
image most recently presented to the display surface is from the
swapchain being destroyed, then either any display resources modified by
presenting images from any swapchain associated with the display surface
**must** be reverted by the implementation to their state prior to the
first present performed on one of these swapchains, or such resources
**must** be left in their current state.

If `swapchain` has exclusive full-screen access, it is released before
the swapchain is destroyed.

Valid Usage

- <a href="#VUID-vkDestroySwapchainKHR-swapchain-01282"
  id="VUID-vkDestroySwapchainKHR-swapchain-01282"></a>
  VUID-vkDestroySwapchainKHR-swapchain-01282  
  All uses of presentable images acquired from `swapchain` **must** have
  completed execution

- <a href="#VUID-vkDestroySwapchainKHR-swapchain-01283"
  id="VUID-vkDestroySwapchainKHR-swapchain-01283"></a>
  VUID-vkDestroySwapchainKHR-swapchain-01283  
  If `VkAllocationCallbacks` were provided when `swapchain` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroySwapchainKHR-swapchain-01284"
  id="VUID-vkDestroySwapchainKHR-swapchain-01284"></a>
  VUID-vkDestroySwapchainKHR-swapchain-01284  
  If no `VkAllocationCallbacks` were provided when `swapchain` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroySwapchainKHR-device-parameter"
  id="VUID-vkDestroySwapchainKHR-device-parameter"></a>
  VUID-vkDestroySwapchainKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroySwapchainKHR-swapchain-parameter"
  id="VUID-vkDestroySwapchainKHR-swapchain-parameter"></a>
  VUID-vkDestroySwapchainKHR-swapchain-parameter  
  If `swapchain` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkDestroySwapchainKHR-pAllocator-parameter"
  id="VUID-vkDestroySwapchainKHR-pAllocator-parameter"></a>
  VUID-vkDestroySwapchainKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroySwapchainKHR-swapchain-parent"
  id="VUID-vkDestroySwapchainKHR-swapchain-parent"></a>
  VUID-vkDestroySwapchainKHR-swapchain-parent  
  If `swapchain` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroySwapchainKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
