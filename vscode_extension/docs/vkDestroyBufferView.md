# vkDestroyBufferView(3) Manual Page

## Name

vkDestroyBufferView - Destroy a buffer view object



## [](#_c_specification)C Specification

To destroy a buffer view, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyBufferView(
    VkDevice                                    device,
    VkBufferView                                bufferView,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the buffer view.
- `bufferView` is the buffer view to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyBufferView-bufferView-00936)VUID-vkDestroyBufferView-bufferView-00936  
  All submitted commands that refer to `bufferView` **must** have completed execution
- [](#VUID-vkDestroyBufferView-bufferView-00937)VUID-vkDestroyBufferView-bufferView-00937  
  If `VkAllocationCallbacks` were provided when `bufferView` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyBufferView-bufferView-00938)VUID-vkDestroyBufferView-bufferView-00938  
  If no `VkAllocationCallbacks` were provided when `bufferView` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyBufferView-device-parameter)VUID-vkDestroyBufferView-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyBufferView-bufferView-parameter)VUID-vkDestroyBufferView-bufferView-parameter  
  If `bufferView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `bufferView` **must** be a valid [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) handle
- [](#VUID-vkDestroyBufferView-pAllocator-parameter)VUID-vkDestroyBufferView-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyBufferView-bufferView-parent)VUID-vkDestroyBufferView-bufferView-parent  
  If `bufferView` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `bufferView` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyBufferView).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0