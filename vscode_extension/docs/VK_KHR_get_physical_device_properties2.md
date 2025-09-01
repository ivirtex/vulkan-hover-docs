# VK\_KHR\_get\_physical\_device\_properties2(3) Manual Page

## Name

VK\_KHR\_get\_physical\_device\_properties2 - instance extension



## [](#_registered_extension_number)Registered Extension Number

60

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_get_physical_device_properties2%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_get_physical_device_properties2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-09-05

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA
- Ian Elliott, Google

## [](#_description)Description

This extension provides new queries for device features, device properties, and format properties that can be easily extended by other extensions, without introducing any further queries. The Vulkan 1.0 feature/limit/formatproperty structures do not include `sType`/`pNext` members. This extension wraps them in new structures with `sType`/`pNext` members, so an application can query a chain of feature/limit/formatproperty structures by constructing the chain and letting the implementation fill them in. A new command is added for each `vkGetPhysicalDevice*` command in core Vulkan 1.0. The new feature structure (and a `pNext` chain of extending structures) can also be passed in to device creation to enable features.

This extension also allows applications to use the physical-device components of device extensions before [vkCreateDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDevice.html) is called.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceFeatures2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2KHR.html)
- [vkGetPhysicalDeviceFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2KHR.html)
- [vkGetPhysicalDeviceImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2KHR.html)
- [vkGetPhysicalDeviceMemoryProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceMemoryProperties2KHR.html)
- [vkGetPhysicalDeviceProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2KHR.html)
- [vkGetPhysicalDeviceQueueFamilyProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2KHR.html)
- [vkGetPhysicalDeviceSparseImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html)

## [](#_new_structures)New Structures

- [VkFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2KHR.html)
- [VkImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2KHR.html)
- [VkPhysicalDeviceImageFormatInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2KHR.html)
- [VkPhysicalDeviceMemoryProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties2KHR.html)
- [VkPhysicalDeviceProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2KHR.html)
- [VkPhysicalDeviceSparseImageFormatInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSparseImageFormatInfo2KHR.html)
- [VkQueueFamilyProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2KHR.html)
- [VkSparseImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatProperties2KHR.html)
- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceFeatures2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_EXTENSION_NAME`
- `VK_KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR`
  - `VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR`

## [](#_examples)Examples

```c++
    // Get features with a hypothetical future extension.
    VkHypotheticalExtensionFeaturesKHR hypotheticalFeatures =
    {
        .sType = VK_STRUCTURE_TYPE_HYPOTHETICAL_FEATURES_KHR,
        .pNext = NULL,
    };

    VkPhysicalDeviceFeatures2KHR features =
    {
        .sType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR,
        .pNext = &hypotheticalFeatures,
    };

    // After this call, features and hypotheticalFeatures have been filled out.
    vkGetPhysicalDeviceFeatures2KHR(physicalDevice, &features);

    // Properties/limits can be chained and queried similarly.

    // Enable some features:
    VkHypotheticalExtensionFeaturesKHR enabledHypotheticalFeatures =
    {
        .sType = VK_STRUCTURE_TYPE_HYPOTHETICAL_FEATURES_KHR,
        .pNext = NULL,
    };

    VkPhysicalDeviceFeatures2KHR enabledFeatures =
    {
        .sType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR,
        .pNext = &enabledHypotheticalFeatures,
    };

    enabledFeatures.features.xyz = VK_TRUE;
    enabledHypotheticalFeatures.abc = VK_TRUE;

    VkDeviceCreateInfo deviceCreateInfo =
    {
        .sType = VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO,
        .pNext = &enabledFeatures,
        ...
        .pEnabledFeatures = NULL,
    };

    VkDevice device;
    vkCreateDevice(physicalDevice, &deviceCreateInfo, NULL, &device);
```

## [](#_version_history)Version History

- Revision 1, 2016-09-12 (Jeff Bolz)
  
  - Internal revisions
- Revision 2, 2016-11-02 (Ian Elliott)
  
  - Added ability for applications to use the physical-device components of device extensions before vkCreateDevice is called.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_get_physical_device_properties2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0