# VkD3D12FenceSubmitInfoKHR(3) Manual Page

## Name

VkD3D12FenceSubmitInfoKHR - Structure specifying values for Direct3D 12 fence-backed semaphores



## [](#_c_specification)C Specification

To specify the values to use when waiting for and signaling semaphores whose [current payload](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing) refers to a Direct3D 12 fence, add a [VkD3D12FenceSubmitInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkD3D12FenceSubmitInfoKHR.html) structure to the `pNext` chain of the [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html) structure. The `VkD3D12FenceSubmitInfoKHR` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `waitSemaphoreValuesCount` is the number of semaphore wait values specified in `pWaitSemaphoreValues`.
- `pWaitSemaphoreValues` is a pointer to an array of `waitSemaphoreValuesCount` values for the corresponding semaphores in [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pWaitSemaphores` to wait for.
- `signalSemaphoreValuesCount` is the number of semaphore signal values specified in `pSignalSemaphoreValues`.
- `pSignalSemaphoreValues` is a pointer to an array of `signalSemaphoreValuesCount` values for the corresponding semaphores in [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pSignalSemaphores` to set when signaled.

## [](#_description)Description

If the semaphore in [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pWaitSemaphores` or [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html)::`pSignalSemaphores` corresponding to an entry in `pWaitSemaphoreValues` or `pSignalSemaphoreValues` respectively does not currently have a [payload](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-payloads) referring to a Direct3D 12 fence, the implementation **must** ignore the value in the `pWaitSemaphoreValues` or `pSignalSemaphoreValues` entry.

Note

As the introduction of the external semaphore handle type `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT` predates that of timeline semaphores, support for importing semaphore payloads from external handles of that type into semaphores created (implicitly or explicitly) with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY` is preserved for backwards compatibility. However, applications **should** prefer importing such handle types into semaphores created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_TIMELINE`, and use the [VkTimelineSemaphoreSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTimelineSemaphoreSubmitInfo.html) structure instead of the `VkD3D12FenceSubmitInfoKHR` structure to specify the values to use when waiting for and signaling such semaphores.

Valid Usage

- [](#VUID-VkD3D12FenceSubmitInfoKHR-waitSemaphoreValuesCount-00079)VUID-VkD3D12FenceSubmitInfoKHR-waitSemaphoreValuesCount-00079  
  `waitSemaphoreValuesCount` **must** be the same value as `VkSubmitInfo`::`waitSemaphoreCount`, where this structure is in the `pNext` chain of a `VkSubmitInfo` structure
- [](#VUID-VkD3D12FenceSubmitInfoKHR-signalSemaphoreValuesCount-00080)VUID-VkD3D12FenceSubmitInfoKHR-signalSemaphoreValuesCount-00080  
  `signalSemaphoreValuesCount` **must** be the same value as `VkSubmitInfo`::`signalSemaphoreCount`, where this structure is in the `pNext` chain of a `VkSubmitInfo` structure

Valid Usage (Implicit)

- [](#VUID-VkD3D12FenceSubmitInfoKHR-sType-sType)VUID-VkD3D12FenceSubmitInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR`
- [](#VUID-VkD3D12FenceSubmitInfoKHR-pWaitSemaphoreValues-parameter)VUID-VkD3D12FenceSubmitInfoKHR-pWaitSemaphoreValues-parameter  
  If `waitSemaphoreValuesCount` is not `0`, and `pWaitSemaphoreValues` is not `NULL`, `pWaitSemaphoreValues` **must** be a valid pointer to an array of `waitSemaphoreValuesCount` `uint64_t` values
- [](#VUID-VkD3D12FenceSubmitInfoKHR-pSignalSemaphoreValues-parameter)VUID-VkD3D12FenceSubmitInfoKHR-pSignalSemaphoreValues-parameter  
  If `signalSemaphoreValuesCount` is not `0`, and `pSignalSemaphoreValues` is not `NULL`, `pSignalSemaphoreValues` **must** be a valid pointer to an array of `signalSemaphoreValuesCount` `uint64_t` values

## [](#_see_also)See Also

[VK\_KHR\_external\_semaphore\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_win32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkD3D12FenceSubmitInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0