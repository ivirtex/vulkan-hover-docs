# VK\_KHR\_device\_group\_creation(3) Manual Page

## Name

VK\_KHR\_device\_group\_creation - instance extension



## [](#_registered_extension_number)Registered Extension Number

71

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_device_group_creation%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_device_group_creation%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-19

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension provides instance-level commands to enumerate groups of physical devices, and to create a logical device from a subset of one of those groups. Such a logical device can then be used with new features in the `VK_KHR_device_group` extension.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkEnumeratePhysicalDeviceGroupsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumeratePhysicalDeviceGroupsKHR.html)

## [](#_new_structures)New Structures

- [VkPhysicalDeviceGroupPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceGroupPropertiesKHR.html)
- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkDeviceGroupDeviceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupDeviceCreateInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_DEVICE_GROUP_CREATION_EXTENSION_NAME`
- `VK_KHR_DEVICE_GROUP_CREATION_SPEC_VERSION`
- `VK_MAX_DEVICE_GROUP_SIZE_KHR`
- Extending [VkMemoryHeapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeapFlagBits.html):
  
  - `VK_MEMORY_HEAP_MULTI_INSTANCE_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR`

## [](#_examples)Examples

```c++
    VkDeviceCreateInfo devCreateInfo = { VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO };
    // (not shown) fill out devCreateInfo as usual.
    uint32_t deviceGroupCount = 0;
    VkPhysicalDeviceGroupPropertiesKHR *props = NULL;

    // Query the number of device groups
    vkEnumeratePhysicalDeviceGroupsKHR(g_vkInstance, &deviceGroupCount, NULL);

    // Allocate and initialize structures to query the device groups
    props = (VkPhysicalDeviceGroupPropertiesKHR *)malloc(deviceGroupCount*sizeof(VkPhysicalDeviceGroupPropertiesKHR));
    for (i = 0; i < deviceGroupCount; ++i) {
        props[i].sType = VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR;
        props[i].pNext = NULL;
    }
    vkEnumeratePhysicalDeviceGroupsKHR(g_vkInstance, &deviceGroupCount, props);

    // If the first device group has more than one physical device. create
    // a logical device using all of the physical devices.
    VkDeviceGroupDeviceCreateInfoKHR deviceGroupInfo = { VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR };
    if (props[0].physicalDeviceCount > 1) {
        deviceGroupInfo.physicalDeviceCount = props[0].physicalDeviceCount;
        deviceGroupInfo.pPhysicalDevices = props[0].physicalDevices;
        devCreateInfo.pNext = &deviceGroupInfo;
    }

    vkCreateDevice(props[0].physicalDevices[0], &devCreateInfo, NULL, &g_vkDevice);
    free(props);
```

## [](#_version_history)Version History

- Revision 1, 2016-10-19 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_device_group_creation).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0