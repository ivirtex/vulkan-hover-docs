# VkExportMetalDeviceInfoEXT(3) Manual Page

## Name

VkExportMetalDeviceInfoEXT - Structure that identifies a VkDevice object and corresponding Metal MTLDevice object



## [](#_c_specification)C Specification

To export the Metal `MTLDevice` object underlying the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) associated with a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) object, include a `VkExportMetalDeviceInfoEXT` structure in the `pNext` chain of the `pMetalObjectsInfo` parameter of a [vkExportMetalObjectsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkExportMetalObjectsEXT.html) call.

The `VkExportMetalDeviceInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_metal_objects
typedef struct VkExportMetalDeviceInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    MTLDevice_id       mtlDevice;
} VkExportMetalDeviceInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `mtlDevice` is the Metal `id<MTLDevice>` object underlying the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) associated with the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) object identified in the call. The implementation will return the `MTLDevice` in this member, or it will return `NULL` if no `MTLDevice` could be found underlying the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExportMetalDeviceInfoEXT-sType-sType)VUID-VkExportMetalDeviceInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_METAL_DEVICE_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMetalDeviceInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0