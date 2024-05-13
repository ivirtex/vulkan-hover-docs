# VkPhysicalDeviceProtectedMemoryProperties(3) Manual Page

## Name

VkPhysicalDeviceProtectedMemoryProperties - Structure describing
protected memory properties that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceProtectedMemoryProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceProtectedMemoryProperties {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           protectedNoFault;
} VkPhysicalDeviceProtectedMemoryProperties;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-limits-protectedNoFault"></span>
  `protectedNoFault` specifies how an implementation behaves when an
  application attempts to write to unprotected memory in a protected
  queue operation, read from protected memory in an unprotected queue
  operation, or perform a query in a protected queue operation. If this
  limit is `VK_TRUE`, such writes will be discarded or have undefined
  values written, reads and queries will return undefined values. If
  this limit is `VK_FALSE`, applications **must** not perform these
  operations. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-protected-access-rules"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-protected-access-rules</a>
  for more information.

If the `VkPhysicalDeviceProtectedMemoryProperties` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceProtectedMemoryProperties-sType-sType"
  id="VUID-VkPhysicalDeviceProtectedMemoryProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceProtectedMemoryProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROTECTED_MEMORY_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceProtectedMemoryProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
