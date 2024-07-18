# VK_KHR_get_physical_device_properties2(3) Manual Page

## Name

VK_KHR_get_physical_device_properties2 - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

60

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_get_physical_device_properties2%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_get_physical_device_properties2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

- Ian Elliott, Google

## <a href="#_description" class="anchor"></a>Description

This extension provides new queries for device features, device
properties, and format properties that can be easily extended by other
extensions, without introducing any further queries. The Vulkan 1.0
feature/limit/formatproperty structures do not include `sType`/`pNext`
members. This extension wraps them in new structures with
`sType`/`pNext` members, so an application can query a chain of
feature/limit/formatproperty structures by constructing the chain and
letting the implementation fill them in. A new command is added for each
`vkGetPhysicalDevice*` command in core Vulkan 1.0. The new feature
structure (and a `pNext` chain of extending structures) can also be
passed in to device creation to enable features.

This extension also allows applications to use the physical-device
components of device extensions before
[vkCreateDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDevice.html) is called.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPhysicalDeviceFeatures2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2KHR.html)

- [vkGetPhysicalDeviceFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2KHR.html)

- [vkGetPhysicalDeviceImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2KHR.html)

- [vkGetPhysicalDeviceMemoryProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties2KHR.html)

- [vkGetPhysicalDeviceProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2KHR.html)

- [vkGetPhysicalDeviceQueueFamilyProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2KHR.html)

- [vkGetPhysicalDeviceSparseImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties2KHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2KHR.html)

- [VkImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2KHR.html)

- [VkPhysicalDeviceImageFormatInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2KHR.html)

- [VkPhysicalDeviceMemoryProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties2KHR.html)

- [VkPhysicalDeviceProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2KHR.html)

- [VkPhysicalDeviceSparseImageFormatInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseImageFormatInfo2KHR.html)

- [VkQueueFamilyProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2KHR.html)

- [VkSparseImageFormatProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties2KHR.html)

- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFeatures2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2KHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_EXTENSION_NAME`

- `VK_KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2_KHR`

  - `VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR`

  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR`

  - `VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR`

## <a href="#_examples" class="anchor"></a>Examples

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-09-12 (Jeff Bolz)

  - Internal revisions

- Revision 2, 2016-11-02 (Ian Elliott)

  - Added ability for applications to use the physical-device components
    of device extensions before vkCreateDevice is called.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_get_physical_device_properties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
