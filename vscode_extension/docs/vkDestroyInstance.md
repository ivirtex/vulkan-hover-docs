# vkDestroyInstance(3) Manual Page

## Name

vkDestroyInstance - Destroy an instance of Vulkan



## [](#_c_specification)C Specification

To destroy an instance, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyInstance(
    VkInstance                                  instance,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `instance` is the handle of the instance to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Prior to destroying an instance, an application is responsible for destroying/freeing any Vulkan objects with explicit `vkDestroy*` or `vkFree*` commands that were created using that instance, or any [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) object retrieved from it, as the first parameter of the corresponding `vkCreate*` or `vkAllocate*` command.

Valid Usage

- [](#VUID-vkDestroyInstance-instance-00629)VUID-vkDestroyInstance-instance-00629  
  All child objects that were created with `instance` or with a [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) retrieved from it, and that **can** be destroyed or freed, **must** have been destroyed or freed prior to destroying `instance`
- [](#VUID-vkDestroyInstance-instance-00630)VUID-vkDestroyInstance-instance-00630  
  If `VkAllocationCallbacks` were provided when `instance` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyInstance-instance-00631)VUID-vkDestroyInstance-instance-00631  
  If no `VkAllocationCallbacks` were provided when `instance` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyInstance-instance-parameter)VUID-vkDestroyInstance-instance-parameter  
  If `instance` is not `NULL`, `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkDestroyInstance-pAllocator-parameter)VUID-vkDestroyInstance-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure

Host Synchronization

- Host access to `instance` **must** be externally synchronized
- Host access to all `VkPhysicalDevice` objects enumerated from `instance` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyInstance).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0