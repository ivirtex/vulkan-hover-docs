# vkQueuePresentKHR(3) Manual Page

## Name

vkQueuePresentKHR - Queue an image for presentation



## <a href="#_c_specification" class="anchor"></a>C Specification

After queueing all rendering commands and transitioning the image to the
correct layout, to queue an image for presentation, call:

``` c
// Provided by VK_KHR_swapchain
VkResult vkQueuePresentKHR(
    VkQueue                                     queue,
    const VkPresentInfoKHR*                     pPresentInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `queue` is a queue that is capable of presentation to the target
  surface’s platform on the same device as the image’s swapchain.

- `pPresentInfo` is a pointer to a
  [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html) structure specifying
  parameters of the presentation.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>There is no requirement for an application to present images in the
same order that they were acquired - applications can arbitrarily
present any image that is currently acquired.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The origin of the native orientation of the surface coordinate system
is not specified in the Vulkan specification; it depends on the
platform. For most platforms the origin is by default upper-left,
meaning the pixel of the presented <a href="VkImage.html">VkImage</a> at
coordinates (0,0) would appear at the upper left pixel of the platform
surface (assuming <code>VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR</code>,
and the display standing the right way up).</p></td>
</tr>
</tbody>
</table>

The result codes `VK_ERROR_OUT_OF_DATE_KHR` and `VK_SUBOPTIMAL_KHR` have
the same meaning when returned by `vkQueuePresentKHR` as they do when
returned by `vkAcquireNextImageKHR`. If any `swapchain` member of
`pPresentInfo` was created with
`VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT`,
`VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT` will be returned if that
swapchain does not have exclusive full-screen access, possibly for
implementation-specific reasons outside of the application’s control. If
multiple swapchains are presented, the result code is determined by
applying the following rules in order:

- If the device is lost, `VK_ERROR_DEVICE_LOST` is returned.

- If any of the target surfaces are no longer available the error
  `VK_ERROR_SURFACE_LOST_KHR` is returned.

- If any of the presents would have a result of
  `VK_ERROR_OUT_OF_DATE_KHR` if issued separately then
  `VK_ERROR_OUT_OF_DATE_KHR` is returned.

- If any of the presents would have a result of
  `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT` if issued separately
  then `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT` is returned.

- If any of the presents would have a result of `VK_SUBOPTIMAL_KHR` if
  issued separately then `VK_SUBOPTIMAL_KHR` is returned.

- Otherwise `VK_SUCCESS` is returned.

Any writes to memory backing the images referenced by the
`pImageIndices` and `pSwapchains` members of `pPresentInfo`, that are
available before [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html) is
executed, are automatically made visible to the read access performed by
the presentation engine. This automatic visibility operation for an
image happens-after the semaphore signal operation, and happens-before
the presentation engine accesses the image.

Presentation is a read-only operation that will not affect the content
of the presentable images. Upon reacquiring the image and transitioning
it away from the `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR` layout, the contents
will be the same as they were prior to transitioning the image to the
present source layout and presenting it. However, if a mechanism other
than Vulkan is used to modify the platform window associated with the
swapchain, the content of all presentable images in the swapchain
becomes undefined.

Calls to `vkQueuePresentKHR` **may** block, but **must** return in
finite time. The processing of the presentation happens in issue order
with other queue operations, but semaphores **must** be used to ensure
that prior rendering and other commands in the specified queue complete
before the presentation begins. The presentation command itself does not
delay processing of subsequent commands on the queue. However,
presentation requests sent to a particular queue are always performed in
order. Exact presentation timing is controlled by the semantics of the
presentation engine and native platform in use.

If an image is presented to a swapchain created from a display surface,
the mode of the associated display will be updated, if necessary, to
match the mode specified when creating the display surface. The mode
switch and presentation of the specified image will be performed as one
atomic operation.

Queueing an image for presentation defines a set of *queue operations*,
including waiting on the semaphores and submitting a presentation
request to the presentation engine. However, the scope of this set of
queue operations does not include the actual processing of the image by
the presentation engine.

If `vkQueuePresentKHR` fails to enqueue the corresponding set of queue
operations, it **may** return `VK_ERROR_OUT_OF_HOST_MEMORY` or
`VK_ERROR_OUT_OF_DEVICE_MEMORY`. If it does, the implementation **must**
ensure that the state and contents of any resources or synchronization
primitives referenced is unaffected by the call or its failure.

If `vkQueuePresentKHR` fails in such a way that the implementation is
unable to make that guarantee, the implementation **must** return
`VK_ERROR_DEVICE_LOST`.

However, if the presentation request is rejected by the presentation
engine with an error `VK_ERROR_OUT_OF_DATE_KHR`,
`VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`, or
`VK_ERROR_SURFACE_LOST_KHR`, the set of queue operations are still
considered to be enqueued and thus any semaphore wait operation
specified in [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html) will execute when
the corresponding queue operation is complete.

`vkQueuePresentKHR` releases the acquisition of the images referenced by
`imageIndices`. The queue family corresponding to the queue
`vkQueuePresentKHR` is executed on **must** have ownership of the
presented images as defined in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-sharing"
target="_blank" rel="noopener">Resource Sharing</a>. `vkQueuePresentKHR`
does not alter the queue family ownership, but the presented images
**must** not be used again before they have been reacquired using
`vkAcquireNextImageKHR`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The application <strong>can</strong> continue to present any acquired
images from a retired swapchain as long as the swapchain has not entered
a state that causes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html">vkQueuePresentKHR</a> to return
<code>VK_ERROR_OUT_OF_DATE_KHR</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkQueuePresentKHR-pSwapchains-01292"
  id="VUID-vkQueuePresentKHR-pSwapchains-01292"></a>
  VUID-vkQueuePresentKHR-pSwapchains-01292  
  Each element of `pSwapchains` member of `pPresentInfo` **must** be a
  swapchain that is created for a surface for which presentation is
  supported from `queue` as determined using a call to
  `vkGetPhysicalDeviceSurfaceSupportKHR`

- <a href="#VUID-vkQueuePresentKHR-pSwapchains-01293"
  id="VUID-vkQueuePresentKHR-pSwapchains-01293"></a>
  VUID-vkQueuePresentKHR-pSwapchains-01293  
  If more than one member of `pSwapchains` was created from a display
  surface, all display surfaces referenced that refer to the same
  display **must** use the same display mode

- <a href="#VUID-vkQueuePresentKHR-pWaitSemaphores-01294"
  id="VUID-vkQueuePresentKHR-pWaitSemaphores-01294"></a>
  VUID-vkQueuePresentKHR-pWaitSemaphores-01294  
  When a semaphore wait operation referring to a binary semaphore
  defined by the elements of the `pWaitSemaphores` member of
  `pPresentInfo` executes on `queue`, there **must** be no other queues
  waiting on the same semaphore

- <a href="#VUID-vkQueuePresentKHR-pWaitSemaphores-03267"
  id="VUID-vkQueuePresentKHR-pWaitSemaphores-03267"></a>
  VUID-vkQueuePresentKHR-pWaitSemaphores-03267  
  All elements of the `pWaitSemaphores` member of `pPresentInfo`
  **must** be created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_BINARY`

- <a href="#VUID-vkQueuePresentKHR-pWaitSemaphores-03268"
  id="VUID-vkQueuePresentKHR-pWaitSemaphores-03268"></a>
  VUID-vkQueuePresentKHR-pWaitSemaphores-03268  
  All elements of the `pWaitSemaphores` member of `pPresentInfo`
  **must** reference a semaphore signal operation that has been
  submitted for execution and any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operations</a> on
  which it depends **must** have also been submitted for execution

Valid Usage (Implicit)

- <a href="#VUID-vkQueuePresentKHR-queue-parameter"
  id="VUID-vkQueuePresentKHR-queue-parameter"></a>
  VUID-vkQueuePresentKHR-queue-parameter  
  `queue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html) handle

- <a href="#VUID-vkQueuePresentKHR-pPresentInfo-parameter"
  id="VUID-vkQueuePresentKHR-pPresentInfo-parameter"></a>
  VUID-vkQueuePresentKHR-pPresentInfo-parameter  
  `pPresentInfo` **must** be a valid pointer to a valid
  [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html) structure

Host Synchronization

- Host access to `queue` **must** be externally synchronized

- Host access to `pPresentInfo->pWaitSemaphores`\[\] **must** be
  externally synchronized

- Host access to `pPresentInfo->pSwapchains`\[\] **must** be externally
  synchronized

Command Properties

| [Command Buffer Levels](#VkCommandBufferLevel) | [Render Pass Scope](#vkCmdBeginRenderPass) | [Video Coding Scope](#vkCmdBeginVideoCodingKHR) | [Supported Queue Types](#VkQueueFlagBits) | [Command Type](#fundamentals-queueoperation-command-types) |
|------------------------------------------------|--------------------------------------------|-------------------------------------------------|-------------------------------------------|------------------------------------------------------------|
| \-                                             | \-                                         | \-                                              | Any                                       | \-                                                         |

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_SUBOPTIMAL_KHR`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_OUT_OF_DATE_KHR`

- `VK_ERROR_SURFACE_LOST_KHR`

- `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html), [VkQueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueue.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkQueuePresentKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
