# vkCreateSharedSwapchainsKHR(3) Manual Page

## Name

vkCreateSharedSwapchainsKHR - Create multiple swapchains that share
presentable images



## <a href="#_c_specification" class="anchor"></a>C Specification

When the [`VK_KHR_display_swapchain`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display_swapchain.html)
extension is enabled, multiple swapchains that share presentable images
are created by calling:

``` c
// Provided by VK_KHR_display_swapchain
VkResult vkCreateSharedSwapchainsKHR(
    VkDevice                                    device,
    uint32_t                                    swapchainCount,
    const VkSwapchainCreateInfoKHR*             pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkSwapchainKHR*                             pSwapchains);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device to create the swapchains for.

- `swapchainCount` is the number of swapchains to create.

- `pCreateInfos` is a pointer to an array of
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structures
  specifying the parameters of the created swapchains.

- `pAllocator` is the allocator used for host memory allocated for the
  swapchain objects when there is no more specific allocator available
  (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a>).

- `pSwapchains` is a pointer to an array of
  [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) handles in which the created
  swapchain objects will be returned.

## <a href="#_description" class="anchor"></a>Description

`vkCreateSharedSwapchainsKHR` is similar to
[vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html), except that it takes
an array of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)
structures, and returns an array of swapchain objects.

The swapchain creation parameters that affect the properties and number
of presentable images **must** match between all the swapchains. If the
displays used by any of the swapchains do not use the same presentable
image layout or are incompatible in a way that prevents sharing images,
swapchain creation will fail with the result code
`VK_ERROR_INCOMPATIBLE_DISPLAY_KHR`. If any error occurs, no swapchains
will be created. Images presented to multiple swapchains **must** be
re-acquired from all of them before being modified. After destroying one
or more of the swapchains, the remaining swapchains and the presentable
images **can** continue to be used.

Valid Usage (Implicit)

- <a href="#VUID-vkCreateSharedSwapchainsKHR-device-parameter"
  id="VUID-vkCreateSharedSwapchainsKHR-device-parameter"></a>
  VUID-vkCreateSharedSwapchainsKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateSharedSwapchainsKHR-pCreateInfos-parameter"
  id="VUID-vkCreateSharedSwapchainsKHR-pCreateInfos-parameter"></a>
  VUID-vkCreateSharedSwapchainsKHR-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of
  `swapchainCount` valid
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structures

- <a href="#VUID-vkCreateSharedSwapchainsKHR-pAllocator-parameter"
  id="VUID-vkCreateSharedSwapchainsKHR-pAllocator-parameter"></a>
  VUID-vkCreateSharedSwapchainsKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateSharedSwapchainsKHR-pSwapchains-parameter"
  id="VUID-vkCreateSharedSwapchainsKHR-pSwapchains-parameter"></a>
  VUID-vkCreateSharedSwapchainsKHR-pSwapchains-parameter  
  `pSwapchains` **must** be a valid pointer to an array of
  `swapchainCount` [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) handles

- <a href="#VUID-vkCreateSharedSwapchainsKHR-swapchainCount-arraylength"
  id="VUID-vkCreateSharedSwapchainsKHR-swapchainCount-arraylength"></a>
  VUID-vkCreateSharedSwapchainsKHR-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

Host Synchronization

- Host access to `pCreateInfos`\[\].surface **must** be externally
  synchronized

- Host access to `pCreateInfos`\[\].oldSwapchain **must** be externally
  synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INCOMPATIBLE_DISPLAY_KHR`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display_swapchain.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateSharedSwapchainsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
