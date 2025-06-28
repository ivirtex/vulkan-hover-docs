# vkCreateSwapchainKHR(3) Manual Page

## Name

vkCreateSwapchainKHR - Create a swapchain



## [](#_c_specification)C Specification

To create a swapchain, call:

```c++
// Provided by VK_KHR_swapchain
VkResult vkCreateSwapchainKHR(
    VkDevice                                    device,
    const VkSwapchainCreateInfoKHR*             pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSwapchainKHR*                             pSwapchain);
```

## [](#_parameters)Parameters

- `device` is the device to create the swapchain for.
- `pCreateInfo` is a pointer to a [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) structure specifying the parameters of the created swapchain.
- `pAllocator` is the allocator used for host memory allocated for the swapchain object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pSwapchain` is a pointer to a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle in which the created swapchain object will be returned.

## [](#_description)Description

As mentioned above, if `vkCreateSwapchainKHR` succeeds, it will return a handle to a swapchain containing an array of at least `pCreateInfo->minImageCount` presentable images.

While acquired by the application, presentable images **can** be used in any way that equivalent non-presentable images **can** be used. A presentable image is equivalent to a non-presentable image created with the following [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) parameters:

  `VkImageCreateInfo` Field Value

`flags`

`VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT` is set if `VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR` is set

`VK_IMAGE_CREATE_PROTECTED_BIT` is set if `VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR` is set

`VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT` and `VK_IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR` are both set if `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR` is set

all other bits are unset

`imageType`

`VK_IMAGE_TYPE_2D`

`format`

`pCreateInfo->imageFormat`

`extent`

{`pCreateInfo->imageExtent.width`, `pCreateInfo->imageExtent.height`, `1`}

`mipLevels`

1

`arrayLayers`

`pCreateInfo->imageArrayLayers`

`samples`

`VK_SAMPLE_COUNT_1_BIT`

`tiling`

`VK_IMAGE_TILING_OPTIMAL`

`usage`

`pCreateInfo->imageUsage`

`sharingMode`

`pCreateInfo->imageSharingMode`

`queueFamilyIndexCount`

`pCreateInfo->queueFamilyIndexCount`

`pQueueFamilyIndices`

`pCreateInfo->pQueueFamilyIndices`

`initialLayout`

`VK_IMAGE_LAYOUT_UNDEFINED`

The `pCreateInfo->surface` **must** not be destroyed until after the swapchain is destroyed.

If `oldSwapchain` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), and the native window referred to by `pCreateInfo->surface` is already associated with a Vulkan swapchain, `VK_ERROR_NATIVE_WINDOW_IN_USE_KHR` **must** be returned.

If `oldSwapchain` is a valid swapchain and there are outstanding calls to `vkWaitForPresent2KHR`, then `vkCreateSwapchainKHR` **may** block until those calls complete.

If the native window referred to by `pCreateInfo->surface` is already associated with a non-Vulkan graphics API surface, `VK_ERROR_NATIVE_WINDOW_IN_USE_KHR` **must** be returned.

The native window referred to by `pCreateInfo->surface` **must** not become associated with a non-Vulkan graphics API surface before all associated Vulkan swapchains have been destroyed.

`vkCreateSwapchainKHR` will return `VK_ERROR_DEVICE_LOST` if the logical device was lost. The `VkSwapchainKHR` is a child of the `device`, and **must** be destroyed before the `device`. However, `VkSurfaceKHR` is not a child of any `VkDevice` and is not affected by the lost device. After successfully recreating a `VkDevice`, the same `VkSurfaceKHR` **can** be used to create a new `VkSwapchainKHR`, provided the previous one was destroyed.

If the `oldSwapchain` parameter of `pCreateInfo` is a valid swapchain, which has exclusive full-screen access, that access is released from `pCreateInfo->oldSwapchain`. If the command succeeds in this case, the newly created swapchain will automatically acquire exclusive full-screen access from `pCreateInfo->oldSwapchain`.

Note

This implicit transfer is intended to avoid exiting and entering full-screen exclusive mode, which may otherwise cause unwanted visual updates to the display.

In some cases, swapchain creation **may** fail if exclusive full-screen mode is requested for application control, but for some implementation-specific reason exclusive full-screen access is unavailable for the particular combination of parameters provided. If this occurs, `VK_ERROR_INITIALIZATION_FAILED` will be returned.

Note

In particular, it will fail if the `imageExtent` member of `pCreateInfo` does not match the extents of the monitor. Other reasons for failure may include the application not being set as high-dpi aware, or if the physical device and monitor are not compatible in this mode.

If the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) includes a [VkSwapchainPresentBarrierCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentBarrierCreateInfoNV.html) structure, then that structure includes additional swapchain creation parameters specific to the present barrier. Swapchain creation **may** fail if the state of the current system restricts the usage of the present barrier feature [VkSurfaceCapabilitiesPresentBarrierNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentBarrierNV.html), or a swapchain itself does not satisfy all the required conditions. In this scenario `VK_ERROR_INITIALIZATION_FAILED` is returned.

When the [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) in [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) is a display surface, then the [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html) in display surface’s [VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplaySurfaceCreateInfoKHR.html) is associated with a particular [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html). Swapchain creation **may** fail if that [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) is not acquired by the application. In this scenario `VK_ERROR_INITIALIZATION_FAILED` is returned.

Valid Usage (Implicit)

- [](#VUID-vkCreateSwapchainKHR-device-parameter)VUID-vkCreateSwapchainKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateSwapchainKHR-pCreateInfo-parameter)VUID-vkCreateSwapchainKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) structure
- [](#VUID-vkCreateSwapchainKHR-pAllocator-parameter)VUID-vkCreateSwapchainKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateSwapchainKHR-pSwapchain-parameter)VUID-vkCreateSwapchainKHR-pSwapchain-parameter  
  `pSwapchain` **must** be a valid pointer to a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkCreateSwapchainKHR-device-queuecount)VUID-vkCreateSwapchainKHR-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_SURFACE_LOST_KHR`
- `VK_ERROR_NATIVE_WINDOW_IN_USE_KHR`
- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_COMPRESSION_EXHAUSTED_EXT`

## [](#_see_also)See Also

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateSwapchainKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0