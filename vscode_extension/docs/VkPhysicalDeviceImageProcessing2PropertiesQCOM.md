# VkPhysicalDeviceImageProcessing2PropertiesQCOM(3) Manual Page

## Name

VkPhysicalDeviceImageProcessing2PropertiesQCOM - Structure containing image processing2 properties



## [](#_c_specification)C Specification

The `VkPhysicalDeviceImageProcessing2PropertiesQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_image_processing2
typedef struct VkPhysicalDeviceImageProcessing2PropertiesQCOM {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         maxBlockMatchWindow;
} VkPhysicalDeviceImageProcessing2PropertiesQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`maxBlockMatchWindow` is a [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html) describing the largest dimensions (`width` and `height`) that **can** be specified for the block match window.

If the `VkPhysicalDeviceImageProcessing2PropertiesQCOM` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These are properties of the image processing2 information of a physical device.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceImageProcessing2PropertiesQCOM-sType-sType)VUID-VkPhysicalDeviceImageProcessing2PropertiesQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_2_PROPERTIES_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_image\_processing2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_image_processing2.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceImageProcessing2PropertiesQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0