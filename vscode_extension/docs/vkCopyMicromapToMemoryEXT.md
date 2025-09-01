# vkCopyMicromapToMemoryEXT(3) Manual Page

## Name

vkCopyMicromapToMemoryEXT - Serialize a micromap on the host



## [](#_c_specification)C Specification

To copy a micromap to host accessible memory, call:

```c++
// Provided by VK_EXT_opacity_micromap
VkResult vkCopyMicromapToMemoryEXT(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    const VkCopyMicromapToMemoryInfoEXT*        pInfo);
```

## [](#_parameters)Parameters

- `device` is the device which owns `pInfo->src`.
- `deferredOperation` is an optional [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) to [request deferral](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations-requesting) for this command.
- `pInfo` is a pointer to a [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapToMemoryInfoEXT.html) structure defining the copy operation.

## [](#_description)Description

This command fulfills the same task as [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapToMemoryEXT.html) but is executed by the host.

This command produces the same results as [vkCmdCopyMicromapToMemoryEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMicromapToMemoryEXT.html), but writes its result directly to a host pointer, and is executed on the host rather than the device. The output **may** not necessarily be bit-for-bit identical, but it can be equally used by either [vkCmdCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToMicromapEXT.html) or [vkCopyMemoryToMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToMicromapEXT.html).

Valid Usage

- [](#VUID-vkCopyMicromapToMemoryEXT-deferredOperation-03678)VUID-vkCopyMicromapToMemoryEXT-deferredOperation-03678  
  Any previous deferred operation that was associated with `deferredOperation` **must** be complete
- [](#VUID-vkCopyMicromapToMemoryEXT-buffer-07568)VUID-vkCopyMicromapToMemoryEXT-buffer-07568  
  The `buffer` used to create `pInfo->src` **must** be bound to host-visible device memory
- [](#VUID-vkCopyMicromapToMemoryEXT-pInfo-07569)VUID-vkCopyMicromapToMemoryEXT-pInfo-07569  
  `pInfo->dst.hostAddress` **must** be a valid host pointer
- [](#VUID-vkCopyMicromapToMemoryEXT-pInfo-07570)VUID-vkCopyMicromapToMemoryEXT-pInfo-07570  
  `pInfo->dst.hostAddress` **must** be aligned to 16 bytes
- [](#VUID-vkCopyMicromapToMemoryEXT-micromapHostCommands-07571)VUID-vkCopyMicromapToMemoryEXT-micromapHostCommands-07571  
  The [`VkPhysicalDeviceOpacityMicromapFeaturesEXT`::`micromapHostCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromapHostCommands) feature **must** be enabled
- [](#VUID-vkCopyMicromapToMemoryEXT-buffer-07572)VUID-vkCopyMicromapToMemoryEXT-buffer-07572  
  The `buffer` used to create `pInfo->src` **must** be bound to memory that was not allocated with multiple instances

Valid Usage (Implicit)

- [](#VUID-vkCopyMicromapToMemoryEXT-device-parameter)VUID-vkCopyMicromapToMemoryEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parameter)VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parameter  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `deferredOperation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkCopyMicromapToMemoryEXT-pInfo-parameter)VUID-vkCopyMicromapToMemoryEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapToMemoryInfoEXT.html) structure
- [](#VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parent)VUID-vkCopyMicromapToMemoryEXT-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_OPERATION_DEFERRED_KHR`
- `VK_OPERATION_NOT_DEFERRED_KHR`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapToMemoryInfoEXT.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCopyMicromapToMemoryEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0