# vkGetFenceStatus(3) Manual Page

## Name

vkGetFenceStatus - Return the status of a fence



## [](#_c_specification)C Specification

To query the status of a fence from the host, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkGetFenceStatus(
    VkDevice                                    device,
    VkFence                                     fence);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the fence.
- `fence` is the handle of the fence to query.

## [](#_description)Description

Upon success, `vkGetFenceStatus` returns the status of the fence object, with the following return codes:

Table 1. Fence Object Status Codes   Status Meaning

`VK_SUCCESS`

The fence specified by `fence` is signaled.

`VK_NOT_READY`

The fence specified by `fence` is unsignaled.

`VK_ERROR_DEVICE_LOST`

The device has been lost. See [Lost Device](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-lost-device).

If a [queue submission](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission) command is pending execution, then the value returned by this command **may** immediately be out of date.

If the device has been lost (see [Lost Device](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-lost-device)), `vkGetFenceStatus` **may** return any of the above status codes. If the device has been lost and `vkGetFenceStatus` is called repeatedly, it will eventually return either `VK_SUCCESS` or `VK_ERROR_DEVICE_LOST`.

Valid Usage (Implicit)

- [](#VUID-vkGetFenceStatus-device-parameter)VUID-vkGetFenceStatus-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetFenceStatus-fence-parameter)VUID-vkGetFenceStatus-fence-parameter  
  `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-vkGetFenceStatus-fence-parent)VUID-vkGetFenceStatus-fence-parent  
  `fence` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_NOT_READY`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_DEVICE_LOST`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetFenceStatus)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0