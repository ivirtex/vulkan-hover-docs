# VkD3D12FenceSubmitInfoKHR(3) Manual Page

## Name

VkD3D12FenceSubmitInfoKHR - Structure specifying values for Direct3D 12
fence-backed semaphores



## <a href="#_c_specification" class="anchor"></a>C Specification

To specify the values to use when waiting for and signaling semaphores
whose <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
target="_blank" rel="noopener">current payload</a> refers to a Direct3D
12 fence, add a
[VkD3D12FenceSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkD3D12FenceSubmitInfoKHR.html) structure to
the `pNext` chain of the [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html) structure.
The `VkD3D12FenceSubmitInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_semaphore_win32
typedef struct VkD3D12FenceSubmitInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           waitSemaphoreValuesCount;
    const uint64_t*    pWaitSemaphoreValues;
    uint32_t           signalSemaphoreValuesCount;
    const uint64_t*    pSignalSemaphoreValues;
} VkD3D12FenceSubmitInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `waitSemaphoreValuesCount` is the number of semaphore wait values
  specified in `pWaitSemaphoreValues`.

- `pWaitSemaphoreValues` is a pointer to an array of
  `waitSemaphoreValuesCount` values for the corresponding semaphores in
  [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`pWaitSemaphores` to wait for.

- `signalSemaphoreValuesCount` is the number of semaphore signal values
  specified in `pSignalSemaphoreValues`.

- `pSignalSemaphoreValues` is a pointer to an array of
  `signalSemaphoreValuesCount` values for the corresponding semaphores
  in [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`pSignalSemaphores` to set when
  signaled.

## <a href="#_description" class="anchor"></a>Description

If the semaphore in [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`pWaitSemaphores`
or [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html)::`pSignalSemaphores` corresponding
to an entry in `pWaitSemaphoreValues` or `pSignalSemaphoreValues`
respectively does not currently have a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-payloads"
target="_blank" rel="noopener">payload</a> referring to a Direct3D 12
fence, the implementation **must** ignore the value in the
`pWaitSemaphoreValues` or `pSignalSemaphoreValues` entry.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>As the introduction of the external semaphore handle type
<code>VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT</code> predates
that of timeline semaphores, support for importing semaphore payloads
from external handles of that type into semaphores created (implicitly
or explicitly) with a <a href="VkSemaphoreType.html">VkSemaphoreType</a>
of <code>VK_SEMAPHORE_TYPE_BINARY</code> is preserved for backwards
compatibility. However, applications <strong>should</strong> prefer
importing such handle types into semaphores created with a <a
href="VkSemaphoreType.html">VkSemaphoreType</a> of
<code>VK_SEMAPHORE_TYPE_TIMELINE</code>, and use the <a
href="VkTimelineSemaphoreSubmitInfo.html">VkTimelineSemaphoreSubmitInfo</a>
structure instead of the <code>VkD3D12FenceSubmitInfoKHR</code>
structure to specify the values to use when waiting for and signaling
such semaphores.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkD3D12FenceSubmitInfoKHR-waitSemaphoreValuesCount-00079"
  id="VUID-VkD3D12FenceSubmitInfoKHR-waitSemaphoreValuesCount-00079"></a>
  VUID-VkD3D12FenceSubmitInfoKHR-waitSemaphoreValuesCount-00079  
  `waitSemaphoreValuesCount` **must** be the same value as
  `VkSubmitInfo`::`waitSemaphoreCount`, where this structure is in the
  `pNext` chain of a `VkSubmitInfo` structure

- <a
  href="#VUID-VkD3D12FenceSubmitInfoKHR-signalSemaphoreValuesCount-00080"
  id="VUID-VkD3D12FenceSubmitInfoKHR-signalSemaphoreValuesCount-00080"></a>
  VUID-VkD3D12FenceSubmitInfoKHR-signalSemaphoreValuesCount-00080  
  `signalSemaphoreValuesCount` **must** be the same value as
  `VkSubmitInfo`::`signalSemaphoreCount`, where this structure is in the
  `pNext` chain of a `VkSubmitInfo` structure

Valid Usage (Implicit)

- <a href="#VUID-VkD3D12FenceSubmitInfoKHR-sType-sType"
  id="VUID-VkD3D12FenceSubmitInfoKHR-sType-sType"></a>
  VUID-VkD3D12FenceSubmitInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR`

- <a href="#VUID-VkD3D12FenceSubmitInfoKHR-pWaitSemaphoreValues-parameter"
  id="VUID-VkD3D12FenceSubmitInfoKHR-pWaitSemaphoreValues-parameter"></a>
  VUID-VkD3D12FenceSubmitInfoKHR-pWaitSemaphoreValues-parameter  
  If `waitSemaphoreValuesCount` is not `0`, and `pWaitSemaphoreValues`
  is not `NULL`, `pWaitSemaphoreValues` **must** be a valid pointer to
  an array of `waitSemaphoreValuesCount` `uint64_t` values

- <a
  href="#VUID-VkD3D12FenceSubmitInfoKHR-pSignalSemaphoreValues-parameter"
  id="VUID-VkD3D12FenceSubmitInfoKHR-pSignalSemaphoreValues-parameter"></a>
  VUID-VkD3D12FenceSubmitInfoKHR-pSignalSemaphoreValues-parameter  
  If `signalSemaphoreValuesCount` is not `0`, and
  `pSignalSemaphoreValues` is not `NULL`, `pSignalSemaphoreValues`
  **must** be a valid pointer to an array of
  `signalSemaphoreValuesCount` `uint64_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_semaphore_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_win32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkD3D12FenceSubmitInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
