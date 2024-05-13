# VkExportFenceCreateInfo(3) Manual Page

## Name

VkExportFenceCreateInfo - Structure specifying handle types that can be
exported from a fence



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a fence whose payload **can** be exported to external handles,
add a [VkExportFenceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportFenceCreateInfo.html) structure
to the `pNext` chain of the [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateInfo.html)
structure. The `VkExportFenceCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkExportFenceCreateInfo {
    VkStructureType                   sType;
    const void*                       pNext;
    VkExternalFenceHandleTypeFlags    handleTypes;
} VkExportFenceCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_external_fence
typedef VkExportFenceCreateInfo VkExportFenceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleTypes` is a bitmask of
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  specifying one or more fence handle types the application **can**
  export from the resulting fence. The application **can** request
  multiple handle types for the same fence.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkExportFenceCreateInfo-handleTypes-01446"
  id="VUID-VkExportFenceCreateInfo-handleTypes-01446"></a>
  VUID-VkExportFenceCreateInfo-handleTypes-01446  
  The bits in `handleTypes` **must** be supported and compatible, as
  reported by
  [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceProperties.html)

Valid Usage (Implicit)

- <a href="#VUID-VkExportFenceCreateInfo-sType-sType"
  id="VUID-VkExportFenceCreateInfo-sType-sType"></a>
  VUID-VkExportFenceCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO`

- <a href="#VUID-VkExportFenceCreateInfo-handleTypes-parameter"
  id="VUID-VkExportFenceCreateInfo-handleTypes-parameter"></a>
  VUID-VkExportFenceCreateInfo-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of
  [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlagBits.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalFenceHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportFenceCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
