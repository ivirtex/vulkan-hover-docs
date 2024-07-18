# vkCreateSwapchainKHR(3) Manual Page

## Name

vkCreateSwapchainKHR - Create a swapchain



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a swapchain, call:

``` c
// Provided by VK_KHR_swapchain
VkResult vkCreateSwapchainKHR(
    VkDevice                                    device,
    const VkSwapchainCreateInfoKHR*             pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSwapchainKHR*                             pSwapchain);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device to create the swapchain for.

- `pCreateInfo` is a pointer to a
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structure
  specifying the parameters of the created swapchain.

- `pAllocator` is the allocator used for host memory allocated for the
  swapchain object when there is no more specific allocator available
  (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a>).

- `pSwapchain` is a pointer to a [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle in which the created swapchain object will be returned.

## <a href="#_description" class="anchor"></a>Description

As mentioned above, if `vkCreateSwapchainKHR` succeeds, it will return a
handle to a swapchain containing an array of at least
`pCreateInfo->minImageCount` presentable images.

While acquired by the application, presentable images **can** be used in
any way that equivalent non-presentable images **can** be used. A
presentable image is equivalent to a non-presentable image created with
the following [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) parameters:

<table id="swapchain-wsi-image-create-info"
class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th
class="tableblock halign-left valign-top"><code>VkImageCreateInfo</code>
Field</th>
<th class="tableblock halign-left valign-top">Value</th>
</tr>
</thead>
<tbody>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>flags</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT</code>
is set if
<code>VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR</code> is
set</p>
<p><code>VK_IMAGE_CREATE_PROTECTED_BIT</code> is set if
<code>VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR</code> is set</p>
<p><code>VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT</code> and
<code>VK_IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR</code> are both set if
<code>VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR</code> is set</p>
<p>all other bits are unset</p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>imageType</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>VK_IMAGE_TYPE_2D</code></p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>format</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>pCreateInfo-&gt;imageFormat</code></p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>extent</code></p></td>
<td
class="tableblock halign-left valign-top"><p>{<code>pCreateInfo-&gt;imageExtent.width</code>,
<code>pCreateInfo-&gt;imageExtent.height</code>,
<code>1</code>}</p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>mipLevels</code></p></td>
<td class="tableblock halign-left valign-top"><p>1</p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>arrayLayers</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>pCreateInfo-&gt;imageArrayLayers</code></p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>samples</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>VK_SAMPLE_COUNT_1_BIT</code></p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>tiling</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>VK_IMAGE_TILING_OPTIMAL</code></p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>usage</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>pCreateInfo-&gt;imageUsage</code></p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>sharingMode</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>pCreateInfo-&gt;imageSharingMode</code></p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>queueFamilyIndexCount</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>pCreateInfo-&gt;queueFamilyIndexCount</code></p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>pQueueFamilyIndices</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>pCreateInfo-&gt;pQueueFamilyIndices</code></p></td>
</tr>
<tr>
<td
class="tableblock halign-left valign-top"><p><code>initialLayout</code></p></td>
<td
class="tableblock halign-left valign-top"><p><code>VK_IMAGE_LAYOUT_UNDEFINED</code></p></td>
</tr>
</tbody>
</table>

The `pCreateInfo->surface` **must** not be destroyed until after the
swapchain is destroyed.

If `oldSwapchain` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the
native window referred to by `pCreateInfo->surface` is already
associated with a Vulkan swapchain, `VK_ERROR_NATIVE_WINDOW_IN_USE_KHR`
**must** be returned.

If the native window referred to by `pCreateInfo->surface` is already
associated with a non-Vulkan graphics API surface,
`VK_ERROR_NATIVE_WINDOW_IN_USE_KHR` **must** be returned.

The native window referred to by `pCreateInfo->surface` **must** not
become associated with a non-Vulkan graphics API surface before all
associated Vulkan swapchains have been destroyed.

`vkCreateSwapchainKHR` will return `VK_ERROR_DEVICE_LOST` if the logical
device was lost. The `VkSwapchainKHR` is a child of the `device`, and
**must** be destroyed before the `device`. However, `VkSurfaceKHR` is
not a child of any `VkDevice` and is not affected by the lost device.
After successfully recreating a `VkDevice`, the same `VkSurfaceKHR`
**can** be used to create a new `VkSwapchainKHR`, provided the previous
one was destroyed.

If the `oldSwapchain` parameter of `pCreateInfo` is a valid swapchain,
which has exclusive full-screen access, that access is released from
`pCreateInfo->oldSwapchain`. If the command succeeds in this case, the
newly created swapchain will automatically acquire exclusive full-screen
access from `pCreateInfo->oldSwapchain`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This implicit transfer is intended to avoid exiting and entering
full-screen exclusive mode, which may otherwise cause unwanted visual
updates to the display.</p></td>
</tr>
</tbody>
</table>

In some cases, swapchain creation **may** fail if exclusive full-screen
mode is requested for application control, but for some
implementation-specific reason exclusive full-screen access is
unavailable for the particular combination of parameters provided. If
this occurs, `VK_ERROR_INITIALIZATION_FAILED` will be returned.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>In particular, it will fail if the <code>imageExtent</code> member of
<code>pCreateInfo</code> does not match the extents of the monitor.
Other reasons for failure may include the application not being set as
high-dpi aware, or if the physical device and monitor are not compatible
in this mode.</p></td>
</tr>
</tbody>
</table>

If the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) includes a
[VkSwapchainPresentBarrierCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentBarrierCreateInfoNV.html)
structure, then that structure includes additional swapchain creation
parameters specific to the present barrier. Swapchain creation **may**
fail if the state of the current system restricts the usage of the
present barrier feature
[VkSurfaceCapabilitiesPresentBarrierNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesPresentBarrierNV.html),
or a swapchain itself does not satisfy all the required conditions. In
this scenario `VK_ERROR_INITIALIZATION_FAILED` is returned.

When the [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) in
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) is a display
surface, then the [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html) in display
surface’s
[VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateInfoKHR.html) is
associated with a particular [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html).
Swapchain creation **may** fail if that
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) is not acquired by the application. In
this scenario `VK_ERROR_INITIALIZATION_FAILED` is returned.

Valid Usage (Implicit)

- <a href="#VUID-vkCreateSwapchainKHR-device-parameter"
  id="VUID-vkCreateSwapchainKHR-device-parameter"></a>
  VUID-vkCreateSwapchainKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateSwapchainKHR-pCreateInfo-parameter"
  id="VUID-vkCreateSwapchainKHR-pCreateInfo-parameter"></a>
  VUID-vkCreateSwapchainKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structure

- <a href="#VUID-vkCreateSwapchainKHR-pAllocator-parameter"
  id="VUID-vkCreateSwapchainKHR-pAllocator-parameter"></a>
  VUID-vkCreateSwapchainKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateSwapchainKHR-pSwapchain-parameter"
  id="VUID-vkCreateSwapchainKHR-pSwapchain-parameter"></a>
  VUID-vkCreateSwapchainKHR-pSwapchain-parameter  
  `pSwapchain` **must** be a valid pointer to a
  [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html) handle

Host Synchronization

- Host access to `pCreateInfo->surface` **must** be externally
  synchronized

- Host access to `pCreateInfo->oldSwapchain` **must** be externally
  synchronized

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

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateSwapchainKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
