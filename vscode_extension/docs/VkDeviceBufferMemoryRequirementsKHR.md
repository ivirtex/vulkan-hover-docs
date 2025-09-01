# VkDeviceBufferMemoryRequirements(3) Manual Page

## Name

VkDeviceBufferMemoryRequirements - (None)



## [](#_c_specification)C Specification

The `VkDeviceBufferMemoryRequirements` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkDeviceBufferMemoryRequirements {
    VkStructureType              sType;
    const void*                  pNext;
    const VkBufferCreateInfo*    pCreateInfo;
} VkDeviceBufferMemoryRequirements;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance4
typedef VkDeviceBufferMemoryRequirements VkDeviceBufferMemoryRequirementsKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pCreateInfo` is a pointer to a [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) structure containing parameters affecting creation of the buffer to query.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDeviceBufferMemoryRequirements-sType-sType)VUID-VkDeviceBufferMemoryRequirements-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_BUFFER_MEMORY_REQUIREMENTS`
- [](#VUID-VkDeviceBufferMemoryRequirements-pNext-pNext)VUID-VkDeviceBufferMemoryRequirements-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDeviceBufferMemoryRequirements-pCreateInfo-parameter)VUID-VkDeviceBufferMemoryRequirements-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) structure

## [](#_see_also)See Also

[VK\_KHR\_maintenance4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance4.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceBufferMemoryRequirements.html), [vkGetDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceBufferMemoryRequirements.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceBufferMemoryRequirements).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0