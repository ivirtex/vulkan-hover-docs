# vkResetFences(3) Manual Page

## Name

vkResetFences - Resets one or more fence objects



## [](#_c_specification)C Specification

To set the state of fences to unsignaled from the host, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkResetFences(
    VkDevice                                    device,
    uint32_t                                    fenceCount,
    const VkFence*                              pFences);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the fences.
- `fenceCount` is the number of fences to reset.
- `pFences` is a pointer to an array of fence handles to reset.

## [](#_description)Description

If any member of `pFences` currently has its [payload imported](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-importing) with temporary permanence, that fence’s prior permanent payload is first restored. The remaining operations described therefore operate on the restored payload.

When [vkResetFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetFences.html) is executed on the host, it defines a *fence unsignal operation* for each fence, which resets the fence to the unsignaled state.

If any member of `pFences` is already in the unsignaled state when [vkResetFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetFences.html) is executed, then [vkResetFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkResetFences.html) has no effect on that fence.

Valid Usage

- [](#VUID-vkResetFences-pFences-01123)VUID-vkResetFences-pFences-01123  
  Each element of `pFences` **must** not be currently associated with any queue command that has not yet completed execution on that queue

Valid Usage (Implicit)

- [](#VUID-vkResetFences-device-parameter)VUID-vkResetFences-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkResetFences-pFences-parameter)VUID-vkResetFences-pFences-parameter  
  `pFences` **must** be a valid pointer to an array of `fenceCount` valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handles
- [](#VUID-vkResetFences-fenceCount-arraylength)VUID-vkResetFences-fenceCount-arraylength  
  `fenceCount` **must** be greater than `0`
- [](#VUID-vkResetFences-pFences-parent)VUID-vkResetFences-pFences-parent  
  Each element of `pFences` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to each member of `pFences` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkResetFences).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0