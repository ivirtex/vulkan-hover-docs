# vkImportFenceWin32HandleKHR(3) Manual Page

## Name

vkImportFenceWin32HandleKHR - Import a fence from a Windows HANDLE



## [](#_c_specification)C Specification

To import a fence payload from a Windows handle, call:

```c++
// Provided by VK_KHR_external_fence_win32
VkResult vkImportFenceWin32HandleKHR(
    VkDevice                                    device,
    const VkImportFenceWin32HandleInfoKHR*      pImportFenceWin32HandleInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the fence.
- `pImportFenceWin32HandleInfo` is a pointer to a [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceWin32HandleInfoKHR.html) structure specifying the fence and import parameters.

## [](#_description)Description

Importing a fence payload from Windows handles does not transfer ownership of the handle to the Vulkan implementation. For handle types defined as NT handles, the application **must** release ownership using the `CloseHandle` system call when the handle is no longer needed.

Applications **can** import the same fence payload into multiple instances of Vulkan, into the same instance from which it was exported, and multiple times into a given Vulkan instance.

Valid Usage

- [](#VUID-vkImportFenceWin32HandleKHR-fence-04448)VUID-vkImportFenceWin32HandleKHR-fence-04448  
  `fence` **must** not be associated with any queue command that has not yet completed execution on that queue

Valid Usage (Implicit)

- [](#VUID-vkImportFenceWin32HandleKHR-device-parameter)VUID-vkImportFenceWin32HandleKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkImportFenceWin32HandleKHR-pImportFenceWin32HandleInfo-parameter)VUID-vkImportFenceWin32HandleKHR-pImportFenceWin32HandleInfo-parameter  
  `pImportFenceWin32HandleInfo` **must** be a valid pointer to a valid [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceWin32HandleInfoKHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_external\_fence\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence_win32.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceWin32HandleInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkImportFenceWin32HandleKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0