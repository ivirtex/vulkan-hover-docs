# VkPresentModeKHR(3) Manual Page

## Name

VkPresentModeKHR - Presentation mode supported for a surface



## [](#_c_specification)C Specification

Possible values of elements of the [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html)::`pPresentModes` array, indicating the supported presentation modes for a surface, are:

```c++
// Provided by VK_KHR_surface
typedef enum VkPresentModeKHR {
    VK_PRESENT_MODE_IMMEDIATE_KHR = 0,
    VK_PRESENT_MODE_MAILBOX_KHR = 1,
    VK_PRESENT_MODE_FIFO_KHR = 2,
    VK_PRESENT_MODE_FIFO_RELAXED_KHR = 3,
  // Provided by VK_KHR_shared_presentable_image
    VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR = 1000111000,
  // Provided by VK_KHR_shared_presentable_image
    VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR = 1000111001,
  // Provided by VK_EXT_present_mode_fifo_latest_ready
    VK_PRESENT_MODE_FIFO_LATEST_READY_EXT = 1000361000,
} VkPresentModeKHR;
```

## [](#_description)Description

- `VK_PRESENT_MODE_IMMEDIATE_KHR` specifies that the presentation engine does not wait for a vertical blanking period to update the current image, meaning this mode **may** result in visible tearing. No internal queuing of presentation requests is needed, as the requests are applied immediately.
- `VK_PRESENT_MODE_MAILBOX_KHR` specifies that the presentation engine waits for the next vertical blanking period to update the current image. Tearing **cannot** be observed. An internal single-entry queue is used to hold pending presentation requests. If the queue is full when a new presentation request is received, the new request replaces the existing entry, and any images associated with the prior entry become available for reuse by the application. One request is removed from the queue and processed during each vertical blanking period in which the queue is non-empty.
- `VK_PRESENT_MODE_FIFO_KHR` specifies that the presentation engine waits for the next vertical blanking period to update the current image. Tearing **cannot** be observed. An internal queue is used to hold pending presentation requests. New requests are appended to the end of the queue, and one request is removed from the beginning of the queue and processed during each vertical blanking period in which the queue is non-empty. This is the only value of `presentMode` that is **required** to be supported.
- `VK_PRESENT_MODE_FIFO_RELAXED_KHR` specifies that the presentation engine generally waits for the next vertical blanking period to update the current image. If a vertical blanking period has already passed since the last update of the current image then the presentation engine does not wait for another vertical blanking period for the update, meaning this mode **may** result in visible tearing in this case. This mode is useful for reducing visual stutter with an application that will mostly present a new image before the next vertical blanking period, but may occasionally be late, and present a new image just after the next vertical blanking period. An internal queue is used to hold pending presentation requests. New requests are appended to the end of the queue, and one request is removed from the beginning of the queue and processed during or after each vertical blanking period in which the queue is non-empty.
- `VK_PRESENT_MODE_FIFO_LATEST_READY_EXT` specifies that the presentation engine waits for the next vertical blanking period to update the current image. Tearing **cannot** be observed. An internal queue is used to hold pending presentation requests. New requests are appended to the end of the queue. At each vertical blanking period, the presentation engine dequeues all successive requests that are ready to be presented from the beginning of the queue. If using `VK_GOOGLE_display_timing` to provide a target present time, the presentation engine will check the specified time for each image. If the target present time is less-than or equal-to the current time, the presentation engine will dequeue the image and check the next one. The image of the last dequeued request will be presented. The other dequeued requests will be dropped.
- `VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR` specifies that the presentation engine and application have concurrent access to a single image, which is referred to as a *shared presentable image*. The presentation engine is only required to update the current image after a new presentation request is received. Therefore the application **must** make a presentation request whenever an update is required. However, the presentation engine **may** update the current image at any point, meaning this mode **may** result in visible tearing.
- `VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR` specifies that the presentation engine and application have concurrent access to a single image, which is referred to as a *shared presentable image*. The presentation engine periodically updates the current image on its regular refresh cycle. The application is only required to make one initial presentation request, after which the presentation engine **must** update the current image without any need for further presentation requests. The application **can** indicate the image contents have been updated by making a presentation request, but this does not guarantee the timing of when it will be updated. This mode **may** result in visible tearing if rendering to the image is not timed correctly.

The supported [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) of the presentable images of a swapchain created for a surface **may** differ depending on the presentation mode, and can be determined as per the table below:

Table 1. Presentable Image Usage Queries   Presentation mode Image usage flags

`VK_PRESENT_MODE_IMMEDIATE_KHR`

[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`supportedUsageFlags`

`VK_PRESENT_MODE_MAILBOX_KHR`

[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`supportedUsageFlags`

`VK_PRESENT_MODE_FIFO_KHR`

[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`supportedUsageFlags`

`VK_PRESENT_MODE_FIFO_RELAXED_KHR`

[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`supportedUsageFlags`

`VK_PRESENT_MODE_FIFO_LATEST_READY_EXT`

[VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`supportedUsageFlags`

`VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR`

[VkSharedPresentSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html)::`sharedPresentSupportedUsageFlags`

`VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR`

[VkSharedPresentSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSharedPresentSurfaceCapabilitiesKHR.html)::`sharedPresentSupportedUsageFlags`

Note

For reference, the mode indicated by `VK_PRESENT_MODE_FIFO_KHR` is equivalent to the behavior of {wgl|glX|egl}SwapBuffers with a swap interval of 1, while the mode indicated by `VK_PRESENT_MODE_FIFO_RELAXED_KHR` is equivalent to the behavior of {wgl|glX}SwapBuffers with a swap interval of -1 (from the {WGL|GLX}\_EXT\_swap\_control\_tear extensions).

## [](#_see_also)See Also

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html), [VkLatencySurfaceCapabilitiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySurfaceCapabilitiesNV.html), [VkSurfacePresentModeCompatibilityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeCompatibilityEXT.html), [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeEXT.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html), [VkSwapchainPresentModeInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModeInfoEXT.html), [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html), [vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html), [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPresentModeKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0