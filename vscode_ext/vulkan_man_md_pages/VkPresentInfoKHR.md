# VkPresentInfoKHR(3) Manual Page

## Name

VkPresentInfoKHR - Structure describing parameters of a queue
presentation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPresentInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_swapchain
typedef struct VkPresentInfoKHR {
    VkStructureType          sType;
    const void*              pNext;
    uint32_t                 waitSemaphoreCount;
    const VkSemaphore*       pWaitSemaphores;
    uint32_t                 swapchainCount;
    const VkSwapchainKHR*    pSwapchains;
    const uint32_t*          pImageIndices;
    VkResult*                pResults;
} VkPresentInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `waitSemaphoreCount` is the number of semaphores to wait for before
  issuing the present request. The number **may** be zero.

- `pWaitSemaphores` is `NULL` or a pointer to an array of
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) objects with `waitSemaphoreCount`
  entries, and specifies the semaphores to wait for before issuing the
  present request.

- `swapchainCount` is the number of swapchains being presented to by
  this command.

- `pSwapchains` is a pointer to an array of
  [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) objects with `swapchainCount`
  entries.

- `pImageIndices` is a pointer to an array of indices into the array of
  each swapchain’s presentable images, with `swapchainCount` entries.
  Each entry in this array identifies the image to present on the
  corresponding entry in the `pSwapchains` array.

- `pResults` is a pointer to an array of [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html) typed
  elements with `swapchainCount` entries. Applications that do not need
  per-swapchain results **can** use `NULL` for `pResults`. If
  non-`NULL`, each entry in `pResults` will be set to the
  [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html) for presenting the swapchain corresponding
  to the same index in `pSwapchains`.

## <a href="#_description" class="anchor"></a>Description

Before an application **can** present an image, the image’s layout
**must** be transitioned to the `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`
layout, or for a shared presentable image the
`VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR` layout.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>When transitioning the image to
<code>VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR</code> or
<code>VK_IMAGE_LAYOUT_PRESENT_SRC_KHR</code>, there is no need to delay
subsequent processing, or perform any visibility operations (as <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html">vkQueuePresentKHR</a> performs automatic
visibility operations). To achieve this, the <code>dstAccessMask</code>
member of the <a
href="VkImageMemoryBarrier.html">VkImageMemoryBarrier</a>
<strong>should</strong> be set to <code>0</code>, and the
<code>dstStageMask</code> parameter <strong>should</strong> be set to
<code>VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkPresentInfoKHR-pSwapchain-09231"
  id="VUID-VkPresentInfoKHR-pSwapchain-09231"></a>
  VUID-VkPresentInfoKHR-pSwapchain-09231  
  Elements of `pSwapchain` **must** be unique

- <a href="#VUID-VkPresentInfoKHR-pImageIndices-01430"
  id="VUID-VkPresentInfoKHR-pImageIndices-01430"></a>
  VUID-VkPresentInfoKHR-pImageIndices-01430  
  Each element of `pImageIndices` **must** be the index of a presentable
  image acquired from the swapchain specified by the corresponding
  element of the `pSwapchains` array, and the presented image
  subresource **must** be in the `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR` or
  `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR` layout at the time the operation
  is executed on a `VkDevice`

- <a href="#VUID-VkPresentInfoKHR-pNext-06235"
  id="VUID-VkPresentInfoKHR-pNext-06235"></a>
  VUID-VkPresentInfoKHR-pNext-06235  
  If a [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentIdKHR.html) structure is included in
  the `pNext` chain, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-presentId"
  target="_blank" rel="noopener"><code>presentId</code></a> feature is
  not enabled, each `presentIds` entry in that structure **must** be
  NULL

- <a href="#VUID-VkPresentInfoKHR-pSwapchains-09199"
  id="VUID-VkPresentInfoKHR-pSwapchains-09199"></a>
  VUID-VkPresentInfoKHR-pSwapchains-09199  
  If any element of the `pSwapchains` array has been created with
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html),
  all of the elements of this array **must** be created with
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html)

Valid Usage (Implicit)

- <a href="#VUID-VkPresentInfoKHR-sType-sType"
  id="VUID-VkPresentInfoKHR-sType-sType"></a>
  VUID-VkPresentInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PRESENT_INFO_KHR`

- <a href="#VUID-VkPresentInfoKHR-pNext-pNext"
  id="VUID-VkPresentInfoKHR-pNext-pNext"></a>
  VUID-VkPresentInfoKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDeviceGroupPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentInfoKHR.html),
  [VkDisplayPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPresentInfoKHR.html),
  [VkFrameBoundaryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryEXT.html),
  [VkPresentFrameTokenGGP](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentFrameTokenGGP.html),
  [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentIdKHR.html),
  [VkPresentRegionsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentRegionsKHR.html),
  [VkPresentTimesInfoGOOGLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentTimesInfoGOOGLE.html),
  [VkSwapchainPresentFenceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentFenceInfoEXT.html),
  or [VkSwapchainPresentModeInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModeInfoEXT.html)

- <a href="#VUID-VkPresentInfoKHR-sType-unique"
  id="VUID-VkPresentInfoKHR-sType-unique"></a>
  VUID-VkPresentInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPresentInfoKHR-pWaitSemaphores-parameter"
  id="VUID-VkPresentInfoKHR-pWaitSemaphores-parameter"></a>
  VUID-VkPresentInfoKHR-pWaitSemaphores-parameter  
  If `waitSemaphoreCount` is not `0`, `pWaitSemaphores` **must** be a
  valid pointer to an array of `waitSemaphoreCount` valid
  [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handles

- <a href="#VUID-VkPresentInfoKHR-pSwapchains-parameter"
  id="VUID-VkPresentInfoKHR-pSwapchains-parameter"></a>
  VUID-VkPresentInfoKHR-pSwapchains-parameter  
  `pSwapchains` **must** be a valid pointer to an array of
  `swapchainCount` valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) handles

- <a href="#VUID-VkPresentInfoKHR-pImageIndices-parameter"
  id="VUID-VkPresentInfoKHR-pImageIndices-parameter"></a>
  VUID-VkPresentInfoKHR-pImageIndices-parameter  
  `pImageIndices` **must** be a valid pointer to an array of
  `swapchainCount` `uint32_t` values

- <a href="#VUID-VkPresentInfoKHR-pResults-parameter"
  id="VUID-VkPresentInfoKHR-pResults-parameter"></a>
  VUID-VkPresentInfoKHR-pResults-parameter  
  If `pResults` is not `NULL`, `pResults` **must** be a valid pointer to
  an array of `swapchainCount` [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html) values

- <a href="#VUID-VkPresentInfoKHR-swapchainCount-arraylength"
  id="VUID-VkPresentInfoKHR-swapchainCount-arraylength"></a>
  VUID-VkPresentInfoKHR-swapchainCount-arraylength  
  `swapchainCount` **must** be greater than `0`

- <a href="#VUID-VkPresentInfoKHR-commonparent"
  id="VUID-VkPresentInfoKHR-commonparent"></a>
  VUID-VkPresentInfoKHR-commonparent  
  Both of the elements of `pSwapchains`, and the elements of
  `pWaitSemaphores` that are valid handles of non-ignored parameters
  **must** have been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html), [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html),
[vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPresentInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
