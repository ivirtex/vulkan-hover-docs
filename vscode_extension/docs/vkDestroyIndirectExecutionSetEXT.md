# vkDestroyIndirectExecutionSetEXT(3) Manual Page

## Name

vkDestroyIndirectExecutionSetEXT - Destroy an indirect execution set



## [](#_c_specification)C Specification

Destroy an Indirect Execution Set by calling:

```c++
// Provided by VK_EXT_device_generated_commands
void vkDestroyIndirectExecutionSetEXT(
    VkDevice                                    device,
    VkIndirectExecutionSetEXT                   indirectExecutionSet,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the indirect execution set.
- `indirectExecutionSet` is the indirect execution set to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyIndirectExecutionSetEXT-indirectExecutionSet-11025)VUID-vkDestroyIndirectExecutionSetEXT-indirectExecutionSet-11025  
  All submitted commands that refer to `indirectExecutionSet` **must** have completed execution

Valid Usage (Implicit)

- [](#VUID-vkDestroyIndirectExecutionSetEXT-device-parameter)VUID-vkDestroyIndirectExecutionSetEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyIndirectExecutionSetEXT-indirectExecutionSet-parameter)VUID-vkDestroyIndirectExecutionSetEXT-indirectExecutionSet-parameter  
  If `indirectExecutionSet` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `indirectExecutionSet` **must** be a valid [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html) handle
- [](#VUID-vkDestroyIndirectExecutionSetEXT-pAllocator-parameter)VUID-vkDestroyIndirectExecutionSetEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyIndirectExecutionSetEXT-indirectExecutionSet-parent)VUID-vkDestroyIndirectExecutionSetEXT-indirectExecutionSet-parent  
  If `indirectExecutionSet` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `indirectExecutionSet` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyIndirectExecutionSetEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0