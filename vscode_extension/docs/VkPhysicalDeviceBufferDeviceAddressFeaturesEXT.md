# VkPhysicalDeviceBufferDeviceAddressFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceBufferDeviceAddressFeaturesEXT - Structure describing buffer address features that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceBufferDeviceAddressFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_buffer_device_address
typedef struct VkPhysicalDeviceBufferDeviceAddressFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           bufferDeviceAddress;
    VkBool32           bufferDeviceAddressCaptureReplay;
    VkBool32           bufferDeviceAddressMultiDevice;
} VkPhysicalDeviceBufferDeviceAddressFeaturesEXT;
```

```c++
// Provided by VK_EXT_buffer_device_address
typedef VkPhysicalDeviceBufferDeviceAddressFeaturesEXT VkPhysicalDeviceBufferAddressFeaturesEXT;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`bufferDeviceAddress` indicates that the implementation supports accessing buffer memory in shaders as storage buffers via an address queried from [vkGetBufferDeviceAddressEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddressEXT.html).
- []()`bufferDeviceAddressCaptureReplay` indicates that the implementation supports saving and reusing buffer addresses, e.g. for trace capture and replay.
- []()`bufferDeviceAddressMultiDevice` indicates that the implementation supports the [`bufferDeviceAddress`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressEXT) feature for logical devices created with multiple physical devices. If this feature is not supported, buffer addresses **must** not be queried on a logical device created with more than one physical device.

## [](#_description)Description

If the `VkPhysicalDeviceBufferDeviceAddressFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceBufferDeviceAddressFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Note

The `VkPhysicalDeviceBufferDeviceAddressFeaturesEXT` structure has the same members as the `VkPhysicalDeviceBufferDeviceAddressFeatures` structure, but the functionality indicated by the members is expressed differently. The features indicated by the `VkPhysicalDeviceBufferDeviceAddressFeatures` structure requires additional flags to be passed at memory allocation time, and the capture and replay mechanism is built around opaque capture addresses for buffer and memory objects.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceBufferDeviceAddressFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceBufferDeviceAddressFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BUFFER_DEVICE_ADDRESS_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_buffer_device_address.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceBufferDeviceAddressFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0