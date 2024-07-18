# vkWaitForFences(3) Manual Page

## Name

vkWaitForFences - Wait for one or more fences to become signaled



## <a href="#_c_specification" class="anchor"></a>C Specification

To wait for one or more fences to enter the signaled state on the host,
call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkWaitForFences(
    VkDevice                                    device,
    uint32_t                                    fenceCount,
    const VkFence*                              pFences,
    VkBool32                                    waitAll,
    uint64_t                                    timeout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the fences.

- `fenceCount` is the number of fences to wait on.

- `pFences` is a pointer to an array of `fenceCount` fence handles.

- `waitAll` is the condition that **must** be satisfied to successfully
  unblock the wait. If `waitAll` is `VK_TRUE`, then the condition is
  that all fences in `pFences` are signaled. Otherwise, the condition is
  that at least one fence in `pFences` is signaled.

- `timeout` is the timeout period in units of nanoseconds. `timeout` is
  adjusted to the closest value allowed by the implementation-dependent
  timeout accuracy, which **may** be substantially longer than one
  nanosecond, and **may** be longer than the requested period.

## <a href="#_description" class="anchor"></a>Description

If the condition is satisfied when `vkWaitForFences` is called, then
`vkWaitForFences` returns immediately. If the condition is not satisfied
at the time `vkWaitForFences` is called, then `vkWaitForFences` will
block and wait until the condition is satisfied or the `timeout` has
expired, whichever is sooner.

If `timeout` is zero, then `vkWaitForFences` does not wait, but simply
returns the current state of the fences. `VK_TIMEOUT` will be returned
in this case if the condition is not satisfied, even though no actual
wait was performed.

If the condition is satisfied before the `timeout` has expired,
`vkWaitForFences` returns `VK_SUCCESS`. Otherwise, `vkWaitForFences`
returns `VK_TIMEOUT` after the `timeout` has expired.

If device loss occurs (see <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-lost-device"
target="_blank" rel="noopener">Lost Device</a>) before the timeout has
expired, `vkWaitForFences` **must** return in finite time with either
`VK_SUCCESS` or `VK_ERROR_DEVICE_LOST`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>While we guarantee that <code>vkWaitForFences</code>
<strong>must</strong> return in finite time, no guarantees are made that
it returns immediately upon device loss. However, the application can
reasonably expect that the delay will be on the order of seconds and
that calling <code>vkWaitForFences</code> will not result in a
permanently (or seemingly permanently) dead process.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkWaitForFences-device-parameter"
  id="VUID-vkWaitForFences-device-parameter"></a>
  VUID-vkWaitForFences-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkWaitForFences-pFences-parameter"
  id="VUID-vkWaitForFences-pFences-parameter"></a>
  VUID-vkWaitForFences-pFences-parameter  
  `pFences` **must** be a valid pointer to an array of `fenceCount`
  valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handles

- <a href="#VUID-vkWaitForFences-fenceCount-arraylength"
  id="VUID-vkWaitForFences-fenceCount-arraylength"></a>
  VUID-vkWaitForFences-fenceCount-arraylength  
  `fenceCount` **must** be greater than `0`

- <a href="#VUID-vkWaitForFences-pFences-parent"
  id="VUID-vkWaitForFences-pFences-parent"></a>
  VUID-vkWaitForFences-pFences-parent  
  Each element of `pFences` **must** have been created, allocated, or
  retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_TIMEOUT`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkWaitForFences"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
