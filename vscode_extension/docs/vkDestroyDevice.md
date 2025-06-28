# vkDestroyDevice(3) Manual Page

## Name

vkDestroyDevice - Destroy a logical device



## [](#_c_specification)C Specification

To destroy a device, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyDevice(
    VkDevice                                    device,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

To ensure that no work is active on the device, [vkDeviceWaitIdle](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDeviceWaitIdle.html) **can** be used to gate the destruction of the device. Prior to destroying a device, an application is responsible for destroying/freeing any Vulkan objects with explicit `vkDestroy*` or `vkFree*` commands that were created using that device as the first parameter of the corresponding `vkCreate*` or `vkAllocate*` command.

Note

The lifetime of each of these objects is bound by the lifetime of the `VkDevice` object. Therefore, to avoid resource leaks, it is critical that an application explicitly free all of these resources prior to calling `vkDestroyDevice`.

Valid Usage

- [](#VUID-vkDestroyDevice-device-05137)VUID-vkDestroyDevice-device-05137  
  All child objects created on `device` that can be destroyed or freed **must** have been destroyed or freed prior to destroying `device`
- [](#VUID-vkDestroyDevice-device-00379)VUID-vkDestroyDevice-device-00379  
  If `VkAllocationCallbacks` were provided when `device` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyDevice-device-00380)VUID-vkDestroyDevice-device-00380  
  If no `VkAllocationCallbacks` were provided when `device` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyDevice-device-parameter)VUID-vkDestroyDevice-device-parameter  
  If `device` is not `NULL`, `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyDevice-pAllocator-parameter)VUID-vkDestroyDevice-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure

Host Synchronization

- Host access to `device` **must** be externally synchronized
- Host access to all `VkQueue` objects created from `device` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyDevice)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0