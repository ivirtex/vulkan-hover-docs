# vkGetDeferredOperationResultKHR(3) Manual Page

## Name

vkGetDeferredOperationResultKHR - Query the result of a deferred operation



## [](#_c_specification)C Specification

The `vkGetDeferredOperationResultKHR` function is defined as:

```c++
// Provided by VK_KHR_deferred_host_operations
VkResult vkGetDeferredOperationResultKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      operation);
```

## [](#_parameters)Parameters

- `device` is the device which owns `operation`.
- `operation` is the operation whose deferred result is being queried.

## [](#_description)Description

If no command has been deferred on `operation`, `vkGetDeferredOperationResultKHR` returns `VK_SUCCESS`.

If the deferred operation is pending, `vkGetDeferredOperationResultKHR` returns `VK_NOT_READY`.

If the deferred operation is complete, it returns the appropriate return value from the original command. This value **must** be one of the [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html) values which could have been returned by the original command if the operation had not been deferred.

Valid Usage (Implicit)

- [](#VUID-vkGetDeferredOperationResultKHR-device-parameter)VUID-vkGetDeferredOperationResultKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeferredOperationResultKHR-operation-parameter)VUID-vkGetDeferredOperationResultKHR-operation-parameter  
  `operation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkGetDeferredOperationResultKHR-operation-parent)VUID-vkGetDeferredOperationResultKHR-operation-parent  
  `operation` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_NOT_READY`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_deferred\_host\_operations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_deferred_host_operations.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeferredOperationResultKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0