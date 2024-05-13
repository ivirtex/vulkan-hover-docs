# VkExportSemaphoreCreateInfo(3) Manual Page

## Name

VkExportSemaphoreCreateInfo - Structure specifying handle types that can
be exported from a semaphore



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a semaphore whose payload **can** be exported to external
handles, add a
[VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreCreateInfo.html)
structure to the `pNext` chain of the
[VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html) structure. The
`VkExportSemaphoreCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkExportSemaphoreCreateInfo {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalSemaphoreHandleTypeFlags    handleTypes;
} VkExportSemaphoreCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_external_semaphore
typedef VkExportSemaphoreCreateInfo VkExportSemaphoreCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleTypes` is a bitmask of
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  specifying one or more semaphore handle types the application **can**
  export from the resulting semaphore. The application **can** request
  multiple handle types for the same semaphore.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkExportSemaphoreCreateInfo-handleTypes-01124"
  id="VUID-VkExportSemaphoreCreateInfo-handleTypes-01124"></a>
  VUID-VkExportSemaphoreCreateInfo-handleTypes-01124  
  The bits in `handleTypes` **must** be supported and compatible, as
  reported by
  [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html)

Valid Usage (Implicit)

- <a href="#VUID-VkExportSemaphoreCreateInfo-sType-sType"
  id="VUID-VkExportSemaphoreCreateInfo-sType-sType"></a>
  VUID-VkExportSemaphoreCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO`

- <a href="#VUID-VkExportSemaphoreCreateInfo-handleTypes-parameter"
  id="VUID-VkExportSemaphoreCreateInfo-handleTypes-parameter"></a>
  VUID-VkExportSemaphoreCreateInfo-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of
  [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlagBits.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalSemaphoreHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportSemaphoreCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
