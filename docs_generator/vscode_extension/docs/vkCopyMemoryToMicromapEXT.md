# vkCopyMemoryToMicromapEXT(3) Manual Page

## Name

vkCopyMemoryToMicromapEXT - Deserialize a micromap on the host



## [](#_c_specification)C Specification

To copy host accessible memory to a micromap, call:

```c++
// Provided by VK_EXT_opacity_micromap
VkResult vkCopyMemoryToMicromapEXT(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyMemoryToMicromapInfoEXT*        pInfo);
```

## [](#_parameters)Parameters

- `device` is the device which owns `pInfo->dst`.
- `deferredOperation` is an optional [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) to [request deferral](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations-requesting) for this command.
- `pInfo` is a pointer to a [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToMicromapInfoEXT.html) structure defining the copy operation.

## [](#_description)Description

This command fulfills the same task as [vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToMicromapEXT.html) but is executed by the host.

This command can accept micromaps produced by either [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapToMemoryEXT.html) or [vkCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMicromapToMemoryEXT.html).

Valid Usage

- [](#VUID-vkCopyMemoryToMicromapEXT-deferredOperation-03678)VUID-vkCopyMemoryToMicromapEXT-deferredOperation-03678  
  Any previous deferred operation that was associated with `deferredOperation` **must** be complete
- [](#VUID-vkCopyMemoryToMicromapEXT-pInfo-07563)VUID-vkCopyMemoryToMicromapEXT-pInfo-07563  
  `pInfo->src.hostAddress` **must** be a valid host pointer
- [](#VUID-vkCopyMemoryToMicromapEXT-pInfo-07564)VUID-vkCopyMemoryToMicromapEXT-pInfo-07564  
  `pInfo->src.hostAddress` **must** be aligned to 16 bytes
- [](#VUID-vkCopyMemoryToMicromapEXT-buffer-07565)VUID-vkCopyMemoryToMicromapEXT-buffer-07565  
  The `buffer` used to create `pInfo->dst` **must** be bound to host-visible device memory
- [](#VUID-vkCopyMemoryToMicromapEXT-micromapHostCommands-07566)VUID-vkCopyMemoryToMicromapEXT-micromapHostCommands-07566  
  The [`VkPhysicalDeviceOpacityMicromapFeaturesEXT`::`micromapHostCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromapHostCommands) feature **must** be enabled
- [](#VUID-vkCopyMemoryToMicromapEXT-buffer-07567)VUID-vkCopyMemoryToMicromapEXT-buffer-07567  
  The `buffer` used to create `pInfo->dst` **must** be bound to memory that was not allocated with multiple instances

Valid Usage (Implicit)

- [](#VUID-vkCopyMemoryToMicromapEXT-device-parameter)VUID-vkCopyMemoryToMicromapEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parameter)VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parameter  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `deferredOperation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkCopyMemoryToMicromapEXT-pInfo-parameter)VUID-vkCopyMemoryToMicromapEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToMicromapInfoEXT.html) structure
- [](#VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parent)VUID-vkCopyMemoryToMicromapEXT-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_OPERATION_DEFERRED_KHR`
- `VK_OPERATION_NOT_DEFERRED_KHR`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToMicromapInfoEXT.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCopyMemoryToMicromapEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0