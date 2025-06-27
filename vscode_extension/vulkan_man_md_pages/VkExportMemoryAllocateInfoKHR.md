# VkExportMemoryAllocateInfo(3) Manual Page

## Name

VkExportMemoryAllocateInfo - Specify exportable handle types for a
device memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

When allocating memory whose payload **may** be exported to another
process or Vulkan instance, add a
[VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html) structure
to the `pNext` chain of the
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure, specifying
the handle types that **may** be exported.

The [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)
structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkExportMemoryAllocateInfo {
    VkStructureType                    sType;
    const void*                        pNext;
    VkExternalMemoryHandleTypeFlags    handleTypes;
} VkExportMemoryAllocateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_external_memory
typedef VkExportMemoryAllocateInfo VkExportMemoryAllocateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleTypes` is zero or a bitmask of
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  specifying one or more memory handle types the application **can**
  export from the resulting allocation. The application **can** request
  multiple handle types for the same allocation.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkExportMemoryAllocateInfo-handleTypes-00656"
  id="VUID-VkExportMemoryAllocateInfo-handleTypes-00656"></a>
  VUID-VkExportMemoryAllocateInfo-handleTypes-00656  
  The bits in `handleTypes` **must** be supported and compatible, as
  reported by
  [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html)
  or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html)

Valid Usage (Implicit)

- <a href="#VUID-VkExportMemoryAllocateInfo-sType-sType"
  id="VUID-VkExportMemoryAllocateInfo-sType-sType"></a>
  VUID-VkExportMemoryAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO`

- <a href="#VUID-VkExportMemoryAllocateInfo-handleTypes-parameter"
  id="VUID-VkExportMemoryAllocateInfo-handleTypes-parameter"></a>
  VUID-VkExportMemoryAllocateInfo-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMemoryAllocateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
