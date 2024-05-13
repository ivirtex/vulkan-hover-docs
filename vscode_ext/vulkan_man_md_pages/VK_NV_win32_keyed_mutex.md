# VK_NV_win32_keyed_mutex(3) Manual Page

## Name

VK_NV_win32_keyed_mutex - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

59

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_NV_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_win32.html)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to
  [VK_KHR_win32_keyed_mutex](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_win32_keyed_mutex.html) extension

## <a href="#_contact" class="anchor"></a>Contact

- Carsten Rohde <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_win32_keyed_mutex%5D%20@crohde%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_win32_keyed_mutex%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>crohde</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-08-19

**IP Status**  
No known IP claims.

**Contributors**  
- James Jones, NVIDIA

- Carsten Rohde, NVIDIA

## <a href="#_description" class="anchor"></a>Description

Applications that wish to import Direct3D 11 memory objects into the
Vulkan API may wish to use the native keyed mutex mechanism to
synchronize access to the memory between Vulkan and Direct3D. This
extension provides a way for an application to access the keyed mutex
associated with an imported Vulkan memory object when submitting command
buffers to a queue.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html),
  [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html):

  - [VkWin32KeyedMutexAcquireReleaseInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32KeyedMutexAcquireReleaseInfoNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_WIN32_KEYED_MUTEX_EXTENSION_NAME`

- `VK_NV_WIN32_KEYED_MUTEX_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV`

## <a href="#_examples" class="anchor"></a>Examples

``` c
    //
    // Import a memory object from Direct3D 11, and synchronize
    // access to it in Vulkan using keyed mutex objects.
    //

    extern VkPhysicalDevice physicalDevice;
    extern VkDevice device;
    extern HANDLE sharedNtHandle;

    static const VkFormat format = VK_FORMAT_R8G8B8A8_UNORM;
    static const VkExternalMemoryHandleTypeFlagsNV handleType =
        VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_IMAGE_BIT_NV;

    VkPhysicalDeviceMemoryProperties memoryProperties;
    VkExternalImageFormatPropertiesNV properties;
    VkExternalMemoryImageCreateInfoNV externalMemoryImageCreateInfo;
    VkImageCreateInfo imageCreateInfo;
    VkImage image;
    VkMemoryRequirements imageMemoryRequirements;
    uint32_t numMemoryTypes;
    uint32_t memoryType;
    VkImportMemoryWin32HandleInfoNV importMemoryInfo;
    VkMemoryAllocateInfo memoryAllocateInfo;
    VkDeviceMemory mem;
    VkResult result;

    // Figure out how many memory types the device supports
    vkGetPhysicalDeviceMemoryProperties(physicalDevice,
                                        &memoryProperties);
    numMemoryTypes = memoryProperties.memoryTypeCount;

    // Check the external handle type capabilities for the chosen format
    // Importable 2D image support with at least 1 mip level, 1 array
    // layer, and VK_SAMPLE_COUNT_1_BIT using optimal tiling and supporting
    // texturing and color rendering is required.
    result = vkGetPhysicalDeviceExternalImageFormatPropertiesNV(
        physicalDevice,
        format,
        VK_IMAGE_TYPE_2D,
        VK_IMAGE_TILING_OPTIMAL,
        VK_IMAGE_USAGE_SAMPLED_BIT |
        VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT,
        0,
        handleType,
        &properties);

    if ((result != VK_SUCCESS) ||
        !(properties.externalMemoryFeatures &
          VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV)) {
        abort();
    }

    // Set up the external memory image creation info
    memset(&externalMemoryImageCreateInfo,
           0, sizeof(externalMemoryImageCreateInfo));
    externalMemoryImageCreateInfo.sType =
        VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV;
    externalMemoryImageCreateInfo.handleTypes = handleType;
    // Set up the  core image creation info
    memset(&imageCreateInfo, 0, sizeof(imageCreateInfo));
    imageCreateInfo.sType = VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO;
    imageCreateInfo.pNext = &externalMemoryImageCreateInfo;
    imageCreateInfo.format = format;
    imageCreateInfo.extent.width = 64;
    imageCreateInfo.extent.height = 64;
    imageCreateInfo.extent.depth = 1;
    imageCreateInfo.mipLevels = 1;
    imageCreateInfo.arrayLayers = 1;
    imageCreateInfo.samples = VK_SAMPLE_COUNT_1_BIT;
    imageCreateInfo.tiling = VK_IMAGE_TILING_OPTIMAL;
    imageCreateInfo.usage = VK_IMAGE_USAGE_SAMPLED_BIT |
        VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT;
    imageCreateInfo.sharingMode = VK_SHARING_MODE_EXCLUSIVE;
    imageCreateInfo.initialLayout = VK_IMAGE_LAYOUT_UNDEFINED;

    vkCreateImage(device, &imageCreateInfo, NULL, &image);
    vkGetImageMemoryRequirements(device,
                                 image,
                                 &imageMemoryRequirements);

    // For simplicity, just pick the first compatible memory type.
    for (memoryType = 0; memoryType < numMemoryTypes; memoryType++) {
        if ((1 << memoryType) & imageMemoryRequirements.memoryTypeBits) {
            break;
        }
    }

    // At least one memory type must be supported given the prior external
    // handle capability check.
    assert(memoryType < numMemoryTypes);

    // Allocate the external memory object.
    memset(&exportMemoryAllocateInfo, 0, sizeof(exportMemoryAllocateInfo));
    exportMemoryAllocateInfo.sType =
        VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV;
    importMemoryInfo.handleTypes = handleType;
    importMemoryInfo.handle = sharedNtHandle;

    memset(&memoryAllocateInfo, 0, sizeof(memoryAllocateInfo));
    memoryAllocateInfo.sType = VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO;
    memoryAllocateInfo.pNext = &exportMemoryAllocateInfo;
    memoryAllocateInfo.allocationSize = imageMemoryRequirements.size;
    memoryAllocateInfo.memoryTypeIndex = memoryType;

    vkAllocateMemory(device, &memoryAllocateInfo, NULL, &mem);

    vkBindImageMemory(device, image, mem, 0);

    ...

    const uint64_t acquireKey = 1;
    const uint32_t timeout = INFINITE;
    const uint64_t releaseKey = 2;

    VkWin32KeyedMutexAcquireReleaseInfoNV keyedMutex =
        { VK_STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV };
    keyedMutex.acquireCount = 1;
    keyedMutex.pAcquireSyncs = &mem;
    keyedMutex.pAcquireKeys = &acquireKey;
    keyedMutex.pAcquireTimeoutMilliseconds = &timeout;
    keyedMutex.releaseCount = 1;
    keyedMutex.pReleaseSyncs = &mem;
    keyedMutex.pReleaseKeys = &releaseKey;

    VkSubmitInfo submit_info = { VK_STRUCTURE_TYPE_SUBMIT_INFO, &keyedMutex };
    submit_info.commandBufferCount = 1;
    submit_info.pCommandBuffers = &cmd_buf;
    vkQueueSubmit(queue, 1, &submit_info, VK_NULL_HANDLE);
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2016-08-11 (James Jones)

  - Updated sample code based on the NV external memory extensions.

  - Renamed from NVX to NV extension.

  - Added Overview and Description sections.

  - Updated sample code to use the NV external memory extensions.

- Revision 1, 2016-06-14 (Carsten Rohde)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_win32_keyed_mutex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
