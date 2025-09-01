# VkSwapchainKHR(3) Manual Page

## Name

VkSwapchainKHR - Opaque handle to a swapchain object



## [](#_c_specification)C Specification

A swapchain object (a.k.a. swapchain) provides the ability to present rendering results to a surface. Swapchain objects are represented by `VkSwapchainKHR` handles:

```c++
// Provided by VK_KHR_swapchain
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSwapchainKHR)
```

## [](#_description)Description

A swapchain is an abstraction for an array of presentable images that are associated with a surface. The presentable images are represented by `VkImage` objects created by the platform. One image (which **can** be an array image for multiview/stereoscopic-3D surfaces) is displayed at a time, but multiple images **can** be queued for presentation. An application renders to the image, and then queues the image for presentation to the surface.

A native window **cannot** be associated with more than one non-retired swapchain at a time. Further, swapchains **cannot** be created for native windows that have a non-Vulkan graphics API surface associated with them.

Note

The presentation engine is an abstraction for the platform’s compositor or display engine.

The presentation engine **may** be synchronous or asynchronous with respect to the application and/or logical device.

Some implementations **may** use the device’s graphics queue or dedicated presentation hardware to perform presentation.

The presentable images of a swapchain are owned by the presentation engine. An application **can** acquire use of a presentable image from the presentation engine. Use of a presentable image **must** occur only after the image is returned by [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImageKHR.html), and before it is released by [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html). This includes transitioning the image layout and rendering commands.

An application **can** acquire use of a presentable image with [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImageKHR.html). After acquiring a presentable image and before modifying it, the application **must** use a synchronization primitive to ensure that the presentation engine has finished reading from the image. The application **can** then transition the image’s layout, queue rendering commands to it, etc. Finally, the application presents the image with [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html), which releases the acquisition of the image. The application **can** also release the acquisition of the image through [vkReleaseSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseSwapchainImagesKHR.html), if the image is not in use by the device, and skip the present operation.

The presentation engine controls the order in which presentable images are acquired for use by the application.

Note

This allows the platform to handle situations which require out-of-order return of images after presentation. At the same time, it allows the application to generate command buffers referencing all of the images in the swapchain at initialization time, rather than in its main loop.

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireNextImageInfoKHR.html), [VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemorySwapchainInfoKHR.html), [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSwapchainCreateInfoKHR.html), [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html), [VkReleaseSwapchainImagesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseSwapchainImagesInfoKHR.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html), [vkAcquireFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireFullScreenExclusiveModeEXT.html), [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImageKHR.html), [vkCreateSharedSwapchainsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSharedSwapchainsKHR.html), [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html), [vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySwapchainKHR.html), [vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetLatencyTimingsNV.html), [vkGetPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPastPresentationTimingGOOGLE.html), [vkGetRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRefreshCycleDurationGOOGLE.html), [vkGetSwapchainCounterEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSwapchainCounterEXT.html), [vkGetSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSwapchainImagesKHR.html), [vkGetSwapchainStatusKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSwapchainStatusKHR.html), [vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkLatencySleepNV.html), [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html), [vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseFullScreenExclusiveModeEXT.html), [vkSetHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetHdrMetadataEXT.html), [vkSetLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLatencyMarkerNV.html), [vkSetLatencySleepModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLatencySleepModeNV.html), [vkSetLocalDimmingAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLocalDimmingAMD.html), [vkWaitForPresent2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForPresent2KHR.html), [vkWaitForPresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForPresentKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0