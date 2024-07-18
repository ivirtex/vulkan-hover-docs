# VkDeviceBufferMemoryRequirements(3) Manual Page

## Name

VkDeviceBufferMemoryRequirements - (None)



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceBufferMemoryRequirements` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkDeviceBufferMemoryRequirements {
    VkStructureType              sType;
    const void*                  pNext;
    const VkBufferCreateInfo*    pCreateInfo;
} VkDeviceBufferMemoryRequirements;
```

or the equivalent

``` c
// Provided by VK_KHR_maintenance4
typedef VkDeviceBufferMemoryRequirements VkDeviceBufferMemoryRequirementsKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pCreateInfo` is a pointer to a
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structure containing
  parameters affecting creation of the buffer to query.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceBufferMemoryRequirements-sType-sType"
  id="VUID-VkDeviceBufferMemoryRequirements-sType-sType"></a>
  VUID-VkDeviceBufferMemoryRequirements-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_BUFFER_MEMORY_REQUIREMENTS`

- <a href="#VUID-VkDeviceBufferMemoryRequirements-pNext-pNext"
  id="VUID-VkDeviceBufferMemoryRequirements-pNext-pNext"></a>
  VUID-VkDeviceBufferMemoryRequirements-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDeviceBufferMemoryRequirements-pCreateInfo-parameter"
  id="VUID-VkDeviceBufferMemoryRequirements-pCreateInfo-parameter"></a>
  VUID-VkDeviceBufferMemoryRequirements-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance4](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance4.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceBufferMemoryRequirements.html),
[vkGetDeviceBufferMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceBufferMemoryRequirementsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceBufferMemoryRequirements"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
