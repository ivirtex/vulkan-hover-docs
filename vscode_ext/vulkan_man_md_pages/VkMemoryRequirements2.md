# VkMemoryRequirements2(3) Manual Page

## Name

VkMemoryRequirements2 - Structure specifying memory requirements



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryRequirements2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkMemoryRequirements2 {
    VkStructureType         sType;
    void*                   pNext;
    VkMemoryRequirements    memoryRequirements;
} VkMemoryRequirements2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_memory_requirements2, VK_NV_ray_tracing with VK_KHR_get_memory_requirements2 or VK_VERSION_1_1
typedef VkMemoryRequirements2 VkMemoryRequirements2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memoryRequirements` is a
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure describing
  the memory requirements of the resource.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryRequirements2-sType-sType"
  id="VUID-VkMemoryRequirements2-sType-sType"></a>
  VUID-VkMemoryRequirements2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2`

- <a href="#VUID-VkMemoryRequirements2-pNext-pNext"
  id="VUID-VkMemoryRequirements2-pNext-pNext"></a>
  VUID-VkMemoryRequirements2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkMemoryDedicatedRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedRequirements.html)

- <a href="#VUID-VkMemoryRequirements2-sType-unique"
  id="VUID-VkMemoryRequirements2-sType-unique"></a>
  VUID-VkMemoryRequirements2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2.html),
[vkGetBufferMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2KHR.html),
[vkGetDeviceBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceBufferMemoryRequirements.html),
[vkGetDeviceBufferMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceBufferMemoryRequirementsKHR.html),
[vkGetDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageMemoryRequirements.html),
[vkGetDeviceImageMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageMemoryRequirementsKHR.html),
[vkGetGeneratedCommandsMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetGeneratedCommandsMemoryRequirementsNV.html),
[vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html),
[vkGetImageMemoryRequirements2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2KHR.html),
[vkGetPipelineIndirectMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineIndirectMemoryRequirementsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryRequirements2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
