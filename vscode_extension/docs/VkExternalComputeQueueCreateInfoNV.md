# VkExternalComputeQueueCreateInfoNV(3) Manual Page

## Name

VkExternalComputeQueueCreateInfoNV - Structure specifying configuration parameters for external compute queue creation



## [](#_c_specification)C Specification

The `VkExternalComputeQueueCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_external_compute_queue
typedef struct VkExternalComputeQueueCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkQueue            preferredQueue;
} VkExternalComputeQueueCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `preferredQueue` is a `VkQueue` supporting graphics commands.

## [](#_description)Description

When creating a `VkExternalComputeQueueNV`, the `preferredQueue` field is a strong scheduling hint as to which `VkQueue` Vulkan graphics workloads will be submitted to with the expectation that execution will overlap with execution of work submitted by the external API.

Valid Usage (Implicit)

- [](#VUID-VkExternalComputeQueueCreateInfoNV-sType-sType)VUID-VkExternalComputeQueueCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_COMPUTE_QUEUE_CREATE_INFO_NV`
- [](#VUID-VkExternalComputeQueueCreateInfoNV-pNext-pNext)VUID-VkExternalComputeQueueCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkExternalComputeQueueCreateInfoNV-preferredQueue-parameter)VUID-VkExternalComputeQueueCreateInfoNV-preferredQueue-parameter  
  `preferredQueue` **must** be a valid [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html) handle

## [](#_see_also)See Also

[VK\_NV\_external\_compute\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_compute_queue.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateExternalComputeQueueNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateExternalComputeQueueNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalComputeQueueCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0