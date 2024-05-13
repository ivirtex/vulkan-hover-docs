# VkSwapchainKHR(3) Manual Page

## Name

VkSwapchainKHR - Opaque handle to a swapchain object



## <a href="#_c_specification" class="anchor"></a>C Specification

A swapchain object (a.k.a. swapchain) provides the ability to present
rendering results to a surface. Swapchain objects are represented by
`VkSwapchainKHR` handles:

``` c
// Provided by VK_KHR_swapchain
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSwapchainKHR)
```

## <a href="#_description" class="anchor"></a>Description

A swapchain is an abstraction for an array of presentable images that
are associated with a surface. The presentable images are represented by
`VkImage` objects created by the platform. One image (which **can** be
an array image for multiview/stereoscopic-3D surfaces) is displayed at a
time, but multiple images **can** be queued for presentation. An
application renders to the image, and then queues the image for
presentation to the surface.

A native window **cannot** be associated with more than one non-retired
swapchain at a time. Further, swapchains **cannot** be created for
native windows that have a non-Vulkan graphics API surface associated
with them.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The presentation engine is an abstraction for the platform’s
compositor or display engine.</p>
<p>The presentation engine <strong>may</strong> be synchronous or
asynchronous with respect to the application and/or logical device.</p>
<p>Some implementations <strong>may</strong> use the device’s graphics
queue or dedicated presentation hardware to perform
presentation.</p></td>
</tr>
</tbody>
</table>

The presentable images of a swapchain are owned by the presentation
engine. An application **can** acquire use of a presentable image from
the presentation engine. Use of a presentable image **must** occur only
after the image is returned by
[vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html), and before it is
released by [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html). This includes
transitioning the image layout and rendering commands.

An application **can** acquire use of a presentable image with
[vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html). After acquiring a
presentable image and before modifying it, the application **must** use
a synchronization primitive to ensure that the presentation engine has
finished reading from the image. The application **can** then transition
the image’s layout, queue rendering commands to it, etc. Finally, the
application presents the image with
[vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html), which releases the
acquisition of the image. The application **can** also release the
acquisition of the image through
[vkReleaseSwapchainImagesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseSwapchainImagesEXT.html), if the
image is not in use by the device, and skip the present operation.

The presentation engine controls the order in which presentable images
are acquired for use by the application.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This allows the platform to handle situations which require
out-of-order return of images after presentation. At the same time, it
allows the application to generate command buffers referencing all of
the images in the swapchain at initialization time, rather than in its
main loop.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireNextImageInfoKHR.html),
[VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemorySwapchainInfoKHR.html),
[VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSwapchainCreateInfoKHR.html),
[VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html),
[VkReleaseSwapchainImagesInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkReleaseSwapchainImagesInfoEXT.html),
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html),
[vkAcquireFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireFullScreenExclusiveModeEXT.html),
[vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html),
[vkCreateSharedSwapchainsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSharedSwapchainsKHR.html),
[vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html),
[vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySwapchainKHR.html),
[vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetLatencyTimingsNV.html),
[vkGetPastPresentationTimingGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPastPresentationTimingGOOGLE.html),
[vkGetRefreshCycleDurationGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRefreshCycleDurationGOOGLE.html),
[vkGetSwapchainCounterEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSwapchainCounterEXT.html),
[vkGetSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSwapchainImagesKHR.html),
[vkGetSwapchainStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSwapchainStatusKHR.html),
[vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkLatencySleepNV.html),
[vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html),
[vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseFullScreenExclusiveModeEXT.html),
[vkSetHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetHdrMetadataEXT.html),
[vkSetLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLatencyMarkerNV.html),
[vkSetLatencySleepModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLatencySleepModeNV.html),
[vkSetLocalDimmingAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLocalDimmingAMD.html),
[vkWaitForPresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkWaitForPresentKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
